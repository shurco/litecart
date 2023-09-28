package archive

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Archive is ...
type Archive interface {
	Directory(name string) error
	Header(os.FileInfo) (io.Writer, error)
	Close() error
}

func extractFile(path string, mode os.FileMode, data io.Reader, dest string) error {
	target := filepath.Join(dest, filepath.FromSlash(path))
	if !strings.HasPrefix(target, filepath.Clean(dest)+string(os.PathSeparator)) {
		return fmt.Errorf("path %q escapes archive destination", target)
	}

	if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
		return err
	}

	file, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	if _, err := io.Copy(file, data); err != nil {
		file.Close()
		os.Remove(target)
		return err
	}
	return file.Close()
}
