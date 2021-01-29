package main

import (
	"os"
	"path/filepath"
	"strings"
)

// GetExtension extracts the extension from a file path
func GetExtension(path string) string {
	ext := filepath.Ext(path)
	return strings.ToLower(ext)
}

// GetSize returns the size of a file
func GetSize(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return info.Size()
}
