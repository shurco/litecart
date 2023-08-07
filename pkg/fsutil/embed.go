package fsutil

import (
	"embed"
	"io/fs"
	"strings"
)

// EmbedExtract is ...
func EmbedExtract(fileSystem embed.FS, folder string) error {
	return fs.WalkDir(fileSystem, ".", func(a_file_path string, a_file_info fs.DirEntry, parent_err error) error {
		if parent_err != nil {
			return parent_err
		}

		// skip any directories entries, we just want to add files
		if a_file_info.IsDir() {
			return nil
		}

		if strings.HasPrefix(a_file_path, folder) {
			file, err := OpenFile(a_file_path, FsCWFlags, 0666)
			if err != nil {
				return err
			}

			fileContent, err := fileSystem.ReadFile(a_file_path)
			if err != nil {
				return err
			}

			if _, err := WriteOSFile(file, fileContent); err != nil {
				return err
			}
		}

		return nil
	})
}
