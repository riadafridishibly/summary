package main

import (
	"path/filepath"
	"strings"
)

// GetExtension extracts the extension from a file path
func GetExtension(path string) string {
	ext := filepath.Ext(path)
	return strings.ToLower(ext)
}
