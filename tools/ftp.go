package tools

import (
	"context"
	"fmt"
	"github.com/jlaffaye/ftp"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FTP ftp client
type FTP struct {
	conn *ftp.ServerConn
	// 连接ftp信息
	FTPAddress  string
	FTPUsername string
	FTPPassword string
	FTPTimeout  time.Duration
}

func NewFTP(address, username, password string, timeout time.Duration) *FTP {
	return &FTP{
		FTPAddress:  address,
		FTPUsername: username,
		FTPPassword: password,
		FTPTimeout:  timeout,
	}
}

// Connect establishes a connection to an FTP server using the provided context.
//
// It takes a context.Context as a parameter and returns an error.
func (w *FTP) Connect(ctx context.Context) (err error) {
	w.conn, err = ftp.Dial(w.FTPAddress, ftp.DialWithTimeout(w.FTPTimeout))
	if err != nil {
		return fmt.Errorf("ftp dial error:%s", err)
	}
	if err = w.conn.Login(w.FTPUsername, w.FTPPassword); err != nil {
		return fmt.Errorf("ftp login error:%s", err)
	}
	return nil
}

// Close closes the FTP connection.
//
// It returns an error if there was a problem quitting the FTP connection.
func (w *FTP) Close() (err error) {
	if w.conn != nil {
		if err = w.conn.Quit(); err != nil {
			w.conn = nil
			return fmt.Errorf("ftp quit error:%s", err)
		}
	}
	w.conn = nil
	return nil
}

// ChangeDir changes the directory on the FTP server.
//
// It takes a context.Context and a path string as parameters and returns an error.
func (w *FTP) ChangeDir(ctx context.Context, path string) (err error) {
	if w.conn == nil {
		if err = w.Connect(ctx); err != nil {
			return fmt.Errorf("ftp connect error:%s", err)
		}
	}
	if err = w.conn.ChangeDir(path); err != nil {
		return fmt.Errorf("ftp change dir error:%s", err)
	}
	return nil
}

// readBuf reads a buffer from the FTP server.
//
// It takes the following parameters:
// - ctx: the context.Context used for the request.
// - path: the path of the file to read from.
// - offset: the offset from where to start reading.
//
// It returns the buffer read from the FTP server and an error if any.
func (w *FTP) readBuf(ctx context.Context, path string, offset uint64) (buf []byte, err error) {
	var response *ftp.Response
	if response, err = w.conn.RetrFrom(path, offset); err != nil {
		return nil, fmt.Errorf("ftp retr from error:%s", err)
	}
	defer Close(response)
	return io.ReadAll(response)
}

// Download downloads a file from an FTP server and saves it to a specified directory.
//
// Parameters:
//   - ctx: The context.Context used for the download request.
//   - path: The path of the file to download from the FTP server.
//   - saveAsDirectory: The directory where the downloaded file will be saved.
//   - saveAsFilename: (Optional) The name of the file to save as. If not provided, the
//     original filename will be used.
//
// Returns:
// - err: An error if the download or file operations fail.
func (w *FTP) Download(ctx context.Context, path, saveAsDirectory string, saveAsFilename ...string) (err error) {
	if w.conn == nil {
		if err = w.Connect(ctx); err != nil {
			return fmt.Errorf("ftp connect error:%s", err)
		}
	}
	var (
		buf      []byte
		file     *os.File
		filename string
	)

	if len(path) == 0 {
		return fmt.Errorf("path is empty")
	}

	if len(saveAsFilename) == 0 {
		tokens := strings.Split(path, "/")
		filename = tokens[len(tokens)-1]
	} else {
		filename = saveAsFilename[0]
	}

	var fileInfo os.FileInfo
	if fileInfo, err = os.Stat(saveAsDirectory); err != nil {
		if err = os.MkdirAll(saveAsDirectory, os.ModePerm); err != nil {
			return fmt.Errorf("create directory error:%s", err)
		}
	} else if !fileInfo.IsDir() {
		return fmt.Errorf("%s is not a directory", saveAsDirectory)
	}

	saveAs := filepath.Join(saveAsDirectory, filename)

	file, err = os.OpenFile(saveAs, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return fmt.Errorf("open file error:%s", err)
	}
	defer Close(file)

	if buf, err = w.readBuf(ctx, path, 0); err != nil {
		return fmt.Errorf("read buf error:%s", err)
	}

	if _, err = file.Write(buf); err != nil {
		return fmt.Errorf("write file error:%s", err)
	}

	return nil
}

// GetServerConn get ftp server connection
func (w *FTP) GetServerConn() *ftp.ServerConn {
	return w.conn
}
