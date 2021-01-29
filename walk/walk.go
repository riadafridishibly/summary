package walk

import (
	"os"
	"path/filepath"
)

// GetAllFiles returns the absolute path of all files
// under `root` directory
func GetAllFiles(root string) []string {
	filePaths := make([]string, 16)
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		filePaths = append(filePaths, path)
		return nil
	})
	return filePaths
}
