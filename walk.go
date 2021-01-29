package main

import (
	"os"
	"path/filepath"
)

// GetAllFiles returns the path of all files
// under `root` directory
func GetAllFiles(root string) map[string]*FileTypeInfo {
	extCount := make(map[string]*FileTypeInfo)
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Mode().IsRegular() {
			if ext := GetExtension(path); ext != "" {
				if _, ok := extCount[ext]; !ok {
					extCount[ext] = &FileTypeInfo{Ext: ext}
				}
				fileTypeInfo := extCount[ext]
				fileTypeInfo.Count++
				fileTypeInfo.Size += info.Size()
			}
		}
		return nil
	})
	return extCount
}
