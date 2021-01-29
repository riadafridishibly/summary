package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/riadafridishibly/summary/walk"
)

// DataCollector holds a slice of FileTypeInfo
type DataCollector struct {
	ExtTypeInfo []*FileTypeInfo
}

// Populate populates the data into the DataCollector
func (dc *DataCollector) Populate(path string) {
	extToFileTypeInfo := make(map[string]*FileTypeInfo)
	files := walk.GetAllFiles(path)
	for _, file := range files {
		ext := GetExtension(file)
		if ext != "" {
			if v, ok := extToFileTypeInfo[ext]; ok {
				v.Count++
				v.Size += GetSize(file)
			} else {
				extToFileTypeInfo[strings.ToLower(ext)] = InitFileType(file)
			}
		}
	}
	dc.ExtTypeInfo = make([]*FileTypeInfo, len(extToFileTypeInfo))
	i := 0
	for _, v := range extToFileTypeInfo {
		dc.ExtTypeInfo[i] = v
		i++
	}
}

// Print pritns the result in console.
func (dc *DataCollector) Print() {

	// Sort by size [descending...]
	// TODO: add soring criteria...
	sort.Slice(dc.ExtTypeInfo, func(i, j int) bool {
		return dc.ExtTypeInfo[i].Size > dc.ExtTypeInfo[j].Size
	})

	fmt.Printf("Unique file types: %d\n\n", len(dc.ExtTypeInfo))
	fmt.Printf("%15s %6s %10s\n", "Extension", "Count", "Size")
	fmt.Printf("%15s %6s %10s\n", "---------", "-----", "----")
	for _, value := range dc.ExtTypeInfo {
		fmt.Println(value.ToString())
	}
}
