package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

// FileTypeInfo hold information about extension, number of files and size
type FileTypeInfo struct {
	Ext   string
	Count int64
	Size  int64
}

// ToString retuns a string presenting a FileTypeInfo
func (ft FileTypeInfo) ToString() string {
	return fmt.Sprintf("%15s: %6d [%10s]", ft.Ext, ft.Count, humanize.Bytes(uint64(ft.Size)))
}

// InitFileType initialize a fileType object
func InitFileType(path string) *FileTypeInfo {
	return &FileTypeInfo{
		Ext:   GetExtension(path),
		Count: 1,
		Size:  GetSize(path),
	}
}
