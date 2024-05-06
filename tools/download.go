package tools

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/net/gclient"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type DownloadFile struct {
	directory              string
	filename               string
	disableCreateDirectory bool
	timeout                time.Duration
}

func (d *DownloadFile) DisableCreateDirectory() {
	d.disableCreateDirectory = true
}

// SaveAsDirectory save as directory, only directory
func (d *DownloadFile) SaveAsDirectory(directory string) {
	d.directory = directory
}

// SaveAsFilename save as filename, only filename
func (d *DownloadFile) SaveAsFilename(filename string) {
	d.filename = filename
}

// Timeout set timeout
func (d *DownloadFile) Timeout(timeout time.Duration) {
	d.timeout = timeout
}

// DownloadFromUrl download from url
func (d *DownloadFile) DownloadFromUrl(ctx context.Context, url string) (n int64, err error) {

	if d.filename == "" {
		tokens := strings.Split(url, "/")
		d.filename = tokens[len(tokens)-1]
	}

	if d.directory == "" {
		d.directory = "."
	}

	if !d.disableCreateDirectory {
		if err = os.MkdirAll(d.directory, os.ModePerm); err != nil {
			return n, fmt.Errorf("create directory error: %v", err)
		}
	}

	var a os.FileInfo
	if a, err = os.Stat(d.directory); err != nil {
		return n, fmt.Errorf("stat directory error: %v", err)
	}
	if !a.IsDir() {
		return n, fmt.Errorf("directory is not directory")
	}

	var (
		handle   *os.File
		response *gclient.Response
		client   = gclient.New()
		saveAs   = filepath.Join(d.directory, d.filename)
	)

	if d.timeout > 0 {
		client.SetTimeout(d.timeout)
	}

	if response, err = client.Get(ctx, url, nil); err != nil {
		return n, fmt.Errorf("get url error: %v", err)
	}
	defer Close(response.Body)

	if handle, err = os.OpenFile(saveAs, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm); err != nil {
		return n, fmt.Errorf("open file error: %v", err)
	}
	defer Close(handle)

	if n, err = io.Copy(handle, response.Body); err != nil {
		return n, fmt.Errorf("io copy error: %v", err)
	}

	return n, nil
}
