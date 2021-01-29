package main

import (
	"fmt"

	"github.com/riadafridishibly/summary/walk"
)

func main() {
	files := walk.GetAllFiles(".")
	for _, file := range files {
		fmt.Println(file)
	}
}
