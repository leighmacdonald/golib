package golib

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func GetFileSize(fileName string) (int64, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	s, err := f.Stat()
	if err != nil {
		return 0, err
	}
	return s.Size(), nil
}

// FindFile will walk up the directory tree until it find a file. Max depth of 4 or the minRootDir directory
// is matched
func FindFile(p string, minRootDir string) string {
	var dots []string
	for i := 0; i < 4; i++ {
		dir := path.Join(dots...)
		fPath := path.Join(dir, p)
		if Exists(fPath) {
			fp, err := filepath.Abs(fPath)
			if err == nil {
				return fp
			}
			return fp
		}
		if strings.HasSuffix(dir, minRootDir) {
			return p
		}
		dots = append(dots, "..")
	}
	return p
}

func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.Mode().IsDir()
}
