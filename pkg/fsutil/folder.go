package fsutil

import "os"

// IsDir reports whether the named directory exists.
func IsDir(path string) bool {
	if path == "" || len(path) > 468 {
		return false
	}

	if fi, err := os.Stat(path); err == nil {
		return fi.IsDir()
	}
	return false
}

// Workdir get
func Workdir() string {
	dir, _ := os.Getwd()
	return dir
}

// MkDirs batch make multi dirs at once
func MkDirs(perm os.FileMode, dirPaths ...string) error {
	for _, dirPath := range dirPaths {
		if !IsDir(dirPath) {
			if err := os.MkdirAll(dirPath, perm); err != nil {
				return err
			}
		}
	}
	return nil
}

// MkSubDirs batch make multi sub-dirs at once
func MkSubDirs(perm os.FileMode, parentDir string, subDirs ...string) error {
	for _, dirName := range subDirs {
		dirPath := parentDir + "/" + dirName
		if !IsDir(dirPath) {
			if err := os.MkdirAll(dirPath, perm); err != nil {
				return err
			}
		}
	}
	return nil
}
