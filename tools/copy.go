package tools

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// CopyDir copies the contents of the source directories to the destination directory.
//
// Parameters:
// - dest: the destination directory where the contents will be copied to.
// - src: the source directories to copy from.
//
// Returns:
// - error: an error if the directory copy operation fails.
func CopyDir(dest string, src ...string) error {
	dst := strings.Trim(dest, " ")
	if dst == "" {
		return fmt.Errorf("dst is empty")
	}

	var (
		fi  os.FileInfo
		err error
	)

	if fi, err = os.Stat(dst); err != nil {
		return fmt.Errorf("stat %s: %v", dst, err)
	}
	if dst, err = filepath.Abs(dst); err != nil {
		return fmt.Errorf("abs %s: %v", dst, err)
	}
	for _, f := range src {
		if fi, err = os.Stat(f); err != nil {
			return fmt.Errorf("stat %s: %v", f, err)
		}
		if f, err = filepath.Abs(f); err != nil {
			return fmt.Errorf("abs %s: %v", f, err)
		}
		if !fi.IsDir() {
			_, filename := filepath.Split(f)
			dstFile := dst + string(os.PathSeparator) + filename
			return CopyFile(dstFile, f)
		}
		if err = filepath.WalkDir(f, walkDir(dst, f)); err != nil {
			return fmt.Errorf("walk %s: %v", f, err)
		}
	}
	return nil
}

// Copy copies the contents of the file at the given source path to the destination path.
//
// Parameters:
// - dst: the destination path where the file will be copied to.
// - src: the source path of the file to be copied.
//
// Returns:
// - error: an error if the file copy operation fails.
// Deprecated: use CopyDir instead
func Copy(dst string, src ...string) error {
	dst = strings.Trim(dst, " ")
	if dst == "" {
		return fmt.Errorf("dst is empty")
	}

	var (
		fi  os.FileInfo
		err error
	)

	if fi, err = os.Stat(dst); err != nil {
		return fmt.Errorf("stat %s: %v", dst, err)
	}
	if dst, err = filepath.Abs(dst); err != nil {
		return fmt.Errorf("abs %s: %v", dst, err)
	}
	for _, f := range src {
		if fi, err = os.Stat(f); err != nil {
			return fmt.Errorf("stat %s: %v", f, err)
		}
		if f, err = filepath.Abs(f); err != nil {
			return fmt.Errorf("abs %s: %v", f, err)
		}
		if !fi.IsDir() {
			_, filename := filepath.Split(f)
			dstFile := dst + string(os.PathSeparator) + filename
			return CopyFile(dstFile, f)
		}
		if err = filepath.WalkDir(f, walkDir(dst, f)); err != nil {
			return fmt.Errorf("walk %s: %v", f, err)
		}
	}
	return nil
}

// DirHasPrefix checks if the directory path 's' has a prefix 'prefix'.
//
// It takes in three parameters:
// - s (string): The directory path to check.
// - prefix (string): The prefix to compare against the directory path.
// - pathSeparator (rune): The path separator used in the directory path.
//
// It returns a boolean value indicating whether the directory path has the specified prefix.
func DirHasPrefix(s, prefix string) bool {
	if !strings.HasPrefix(s, prefix) {
		return false
	}
	// but
	// src: /path/to/go-common/tools/test_copy
	// dst: /path/to/go-common/tools/test_copy_file
	// is not working
	// must compare directory name every time
	sourceSlice := strings.Split(s, string(os.PathSeparator))
	prefixSlice := strings.Split(prefix, string(os.PathSeparator))
	if len(prefixSlice) > len(sourceSlice) {
		return false
	}
	for i := 0; i < len(prefixSlice); i++ {
		if prefixSlice[i] != sourceSlice[i] {
			return false
		}
	}
	return true
}

func walkDir(rootDst, rootSrc string) fs.WalkDirFunc {
	return func(path string, d fs.DirEntry, err error) error {

		if DirHasPrefix(path, rootDst) {
			return nil
		}

		subPath := strings.ReplaceAll(path, rootSrc, "")
		if subPath == "" {
			return nil
		}
		dstPath := rootDst + subPath
		srcPath := rootSrc + subPath

		if d.IsDir() {

			_ = os.MkdirAll(dstPath, os.ModePerm)
			return nil
		}

		var info os.FileInfo
		if info, err = d.Info(); err != nil {
			return nil
		}
		if (info.Mode()&os.ModeSymlink != 0) || !info.Mode().IsRegular() {
			return nil
		}

		return CopyFile(dstPath, srcPath)
	}
}

// CopyFile copies the contents of the file at the given source path to the destination path.
//
// Parameters:
// - dst: the destination path where the file will be copied to.
// - src: the source path of the file to be copied.
//
// Returns:
// - error: an error if the file copy operation fails.
func CopyFile(dst, src string) error {
	var (
		newHandle, oldHandle *os.File
		err                  error
	)
	//fmt.Println(dst, src)
	if newHandle, err = os.OpenFile(dst, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm); err != nil {
		return fmt.Errorf("open %s: %v", dst, err)
	}
	defer Close(newHandle)

	if oldHandle, err = os.Open(src); err != nil {
		return fmt.Errorf("open %s: %v", src, err)
	}
	defer Close(oldHandle)

	if _, err = io.Copy(newHandle, oldHandle); err != nil {
		return fmt.Errorf("copy %s to %s: %v", src, dst, err)
	}
	return nil
}
