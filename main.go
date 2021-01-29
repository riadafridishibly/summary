package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/riadafridishibly/summary/walk"
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

type fileType struct {
	Ext   string
	Count int64
	Size  int64
}

func (ft fileType) ToString() string {
	return fmt.Sprintf("%15s: %6d [%10s]", ft.Ext, ft.Count, humanize.Bytes(uint64(ft.Size)))
}

// initFileType initialize a fileType object
func initFileType(path string) *fileType {
	return &fileType{
		Ext:   GetExtension(path),
		Count: 1,
		Size:  GetSize(path),
	}
}

func main() {
	extCount := make(map[string]*fileType)

	files := walk.GetAllFiles(".")
	for _, file := range files {
		ext := GetExtension(file)
		if ext != "" {
			if v, ok := extCount[ext]; ok {
				v.Count++
				v.Size += GetSize(file)
			} else {
				extCount[strings.ToLower(ext)] = initFileType(file)
			}
		}
	}

	keys := make([]string, len(extCount))
	i := 0
	for k := range extCount {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	fmt.Printf("Unique file types: %d\n\n", len(keys))
	fmt.Printf("%15s %6s %10s\n", "Extension", "Count", "Size")
	fmt.Printf("%15s %6s %10s\n", "---------", "-----", "----")
	for _, k := range keys {
		value := extCount[k]
		fmt.Println(value.ToString())
	}
}
