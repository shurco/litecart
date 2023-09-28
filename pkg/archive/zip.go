package archive

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

type ZipArchive struct {
	dir  string
	zipw *zip.Writer
	file io.Closer
}

// NewZipArchive is ...
func NewZipArchive(w io.WriteCloser) Archive {
	return &ZipArchive{"", zip.NewWriter(w), w}
}

// Directory is ...
func (a *ZipArchive) Directory(name string) error {
	a.dir = name + "/"
	return nil
}

// Header is ...
func (a *ZipArchive) Header(fi os.FileInfo) (io.Writer, error) {
	head, err := zip.FileInfoHeader(fi)
	if err != nil {
		return nil, fmt.Errorf("can't make zip header: %v", err)
	}
	head.Name = a.dir + head.Name
	head.Method = zip.Deflate
	w, err := a.zipw.CreateHeader(head)
	if err != nil {
		return nil, fmt.Errorf("can't add zip header: %v", err)
	}
	return w, nil
}

// Close is ...
func (a *ZipArchive) Close() error {
	if err := a.zipw.Close(); err != nil {
		return err
	}
	return a.file.Close()
}

// Extract extracts the zip archive at src to dest.
func ExtractZip(src, dest string) error {
	ar, err := os.Open(src)
	if err != nil {
		return err
	}
	defer ar.Close()

	info, err := ar.Stat()
	if err != nil {
		return err
	}
	zr, err := zip.NewReader(ar, info.Size())
	if err != nil {
		return err
	}

	for _, zf := range zr.File {
		if !zf.Mode().IsRegular() {
			continue
		}

		data, err := zf.Open()
		if err != nil {
			return err
		}
		err = extractFile(zf.Name, zf.Mode(), data, dest)
		data.Close()
		if err != nil {
			return fmt.Errorf("extract %s: %v", zf.Name, err)
		}
	}
	return nil
}
