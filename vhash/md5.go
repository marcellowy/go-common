package vhash

// Package tools
// Copyright 2023 marcello<volibearw@gmail.com>. All rights reserved.

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// Md5 calculates the MD5 hash of the input string.
//
// It takes a string parameter and returns the MD5 hash as a string.// Md5 calculates the MD5 hash of the input string.
//
// It takes a string parameter and returns the MD5 hash as a string.
func Md5(s string) string {
	if s == "" {
		return ""
	}
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}

// Md5File calculates the MD5 hash of a file specified by the filePath parameter.
//
// It takes a string filePath parameter and returns the MD5 hash as a string and an error.
func Md5File(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
