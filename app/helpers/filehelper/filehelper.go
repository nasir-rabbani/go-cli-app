package filehelper

import (
	"os"
	"path/filepath"
)

// FilePointer - Returns file pointer for the given path
func FilePointer(path string) (*os.File, error) {
	parentDir := filepath.Dir(path)
	os.MkdirAll(parentDir, os.ModePerm)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return nil, err
	}
	return f, nil
}
