package tools

import (
	"archive/zip"
	"context"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func UnzipGBKHandler(f **zip.File) {
	//try convert to utf-8 charset from gbk
	var err error
	reader := transform.NewReader(strings.NewReader((*f).Name), simplifiedchinese.GBK.NewDecoder())
	var newName []byte
	if newName, err = io.ReadAll(reader); err == nil {
		(*f).Name = string(newName)
	}
}

// Unzip unzip zip file
func Unzip(ctx context.Context, filename string, saveAs string, handler func(file **zip.File)) error {
	archive, err := zip.OpenReader(filename)
	if err != nil {
		return fmt.Errorf("read filename %s err: %v", filename, err)
	}
	defer Close(archive)
	for _, f := range archive.File {
		if handler != nil {
			handler(&f)
		}
		filePath := filepath.Join(saveAs, f.Name)
		if !strings.HasPrefix(filePath, filepath.Clean(saveAs)+string(os.PathSeparator)) {
			return fmt.Errorf("invalid file path: %s file: %s", filePath, f.Name)
		}
		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(filePath, os.ModePerm)
			continue
		}
		if err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return fmt.Errorf("create directory err: %v", err)
		}
		if err = unzipCopy(filePath, f); err != nil {
			return fmt.Errorf("copy file err: %v", err)
		}
	}
	return nil
}

func unzipCopy(filePath string, f *zip.File) (err error) {
	var dstFile *os.File
	dstFile, err = os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return fmt.Errorf("create file %s err: %v", filePath, err)
	}
	defer Close(dstFile)

	var fileInArchive io.ReadCloser
	fileInArchive, err = f.Open()
	if err != nil {
		return fmt.Errorf("open handle err: %v", err)
	}
	defer Close(fileInArchive)

	if _, err = io.Copy(dstFile, fileInArchive); err != nil {
		return fmt.Errorf("io copy err: %v", err)
	}

	return nil
}

// Zip compresses the specified files or dirs to zip archive.
// If a path is a dir don't need to specify the trailing path separator.
// For example calling Zip("archive.zip", "dir", "csv/baz.csv") will get archive.zip and the content of which is
// baz.csv
// dir
// ├── bar.txt
// └── foo.txt
// Note that if a file is a symbolic link it will be skipped.
func Zip(zipPath string, paths ...string) error {
	// Create zip file and it's parent dir.
	if err := os.MkdirAll(filepath.Dir(zipPath), os.ModePerm); err != nil {
		return err
	}
	archive, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer archive.Close()

	// New zip writer.
	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	// Traverse the file or directory.
	for _, rootPath := range paths {
		// Remove the trailing path separator if path is a directory.
		//rootPath = strings.TrimSuffix(rootPath, string(os.PathSeparator))
		// Visit all the files or directories in the tree.
		err = filepath.WalkDir(rootPath, walkZipFunc(rootPath, zipWriter))
		if err != nil {
			return err
		}
	}
	return nil
}

func walkZipFunc(rootPath string, zipWriter *zip.Writer) fs.WalkDirFunc {
	return func(path string, inf fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		var info os.FileInfo
		if info, err = os.Stat(path); err != nil {
			return fmt.Errorf("stat err: %v", err)
		}

		// If a file is a symbolic link it will be skipped.
		if info.Mode()&os.ModeSymlink != 0 {
			return nil
		}

		// Create a local file header.
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Set compression method.
		header.Method = zip.Deflate

		rootPathAbs, err := filepath.Abs(rootPath)
		if err != nil {
			return fmt.Errorf("abs err: %v", err)
		}
		pathAbs, err := filepath.Abs(path)
		if err != nil {
			return fmt.Errorf("abs err: %v", err)
		}
		name := strings.ReplaceAll(pathAbs, rootPathAbs, "")
		if len(name) > 0 {
			name = name[1:]
		}
		if len(name) == 0 {
			return nil
		}
		if os.PathSeparator == '\\' {
			name = strings.ReplaceAll(name, "\\", "/")
		}
		header.Name = name

		// Create writer for the file header and save content of the file.
		headerWriter, err := zipWriter.CreateHeader(header)
		if err != nil {
			return fmt.Errorf("create header err: %v", err)
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("open err: %v", err)
		}
		defer Close(f)
		_, err = io.Copy(headerWriter, f)
		if err != nil {
			return fmt.Errorf("copy err: %v", err)
		}
		return nil
	}
}
