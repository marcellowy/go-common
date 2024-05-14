// Package tools
// Copyright 2023 marcello<volibearw@gmail.com>. All rights reserved.
package tools

import (
	"fmt"
	"os"
	"path/filepath"
)

// PathExists check path
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetCurrentDirectory return current execute file directory
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return dir
}

// RemoveLastSeparator removes the last separator from the given path string.
//
// Parameters:
// - path: the path string from which the last separator needs to be removed.
//
// Returns:
// - string: the modified path string with the last separator removed.
func RemoveLastSeparator(path string) string {
	var length = len(path) - 1
	if length == -1 {
		return ""
	}
	if path[length] == '\\' {
		return path[:length]
	}
	if path[length] == '/' {
		return path[:length]
	}
	return path
}

// ReCreateDirectory removes the directory at the given path and creates a new empty directory at the same path.
//
// Parameters:
// - dir: the path of the directory to be removed and recreated.
//
// Returns:
// - error: an error if any occurred during the removal or creation of the directory.
func ReCreateDirectory(dir string) error {

	var fileInfo, err = os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			// not exist
		} else {
			// other error
			return err
		}
	} else {
		if !fileInfo.IsDir() {
			// exist and is not dir
			return fmt.Errorf("%s is not a directory", dir)
		}

		// exist and is dir
		if err = os.RemoveAll(dir); err != nil {
			return fmt.Errorf("remove %s error: %s", dir, err.Error())
		}
	}

	// create
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("create %s error: %s", dir, err.Error())
	}
	return nil
}

// CreateEmptyFile creates an empty file at the given path.
// if the file already exists, it will clean content
func CreateEmptyFile(filename string) error {
	dir, file := filepath.Split(filename)
	if file == "" {
		return fmt.Errorf("invalid filename: %s", filename)
	}

	var (
		fileInfo os.FileInfo
		err      error
	)

	if fileInfo, err = os.Stat(dir); err == nil {

		if !fileInfo.IsDir() {
			return fmt.Errorf("%s exists and is not a directory", dir)
		}
	} else {
		_ = os.MkdirAll(dir, os.ModePerm)
	}

	return os.WriteFile(filename, []byte{}, os.ModePerm)
}
