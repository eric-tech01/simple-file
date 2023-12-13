package file

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GetCurrentDirectory ...
func GetCurrentDirectory() string {
	if dir, err := os.Getwd(); err == nil {
		return dir
	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("err", err)
	}

	return strings.Replace(dir, "\\", "/", -1)
}

// GetCurrentPackage ...
func GetCurrentPackage() string {
	if dir, err := os.Getwd(); err == nil {
		return dir
	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("err", err)
	}

	return strings.Replace(dir, "\\", "/", -1)
}
