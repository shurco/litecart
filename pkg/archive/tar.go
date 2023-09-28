package archive

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"time"
)

// TarArchive is ...
type TarArchive struct {
	dir  string
	tarw *tar.Writer
	gzw  *gzip.Writer
	file io.Closer
}

// NewTarArchive is ...
func NewTarArchive(w io.WriteCloser) Archive {
	gzw := gzip.NewWriter(w)
	tarw := tar.NewWriter(gzw)
	return &TarArchive{"", tarw, gzw, w}
}

// Directory is ...
func (a *TarArchive) Directory(name string) error {
	a.dir = name + "/"
	return a.tarw.WriteHeader(&tar.Header{
		Name:     a.dir,
		Mode:     0755,
		Typeflag: tar.TypeDir,
		ModTime:  time.Now(),
	})
}

// Header is ...
func (a *TarArchive) Header(fi os.FileInfo) (io.Writer, error) {
	head, err := tar.FileInfoHeader(fi, "")
	if err != nil {
		return nil, fmt.Errorf("can't make tar header: %v", err)
	}
	head.Name = a.dir + head.Name
	if err := a.tarw.WriteHeader(head); err != nil {
		return nil, fmt.Errorf("can't add tar header: %v", err)
	}
	return a.tarw, nil
}

// Close is ...
func (a *TarArchive) Close() error {
	if err := a.tarw.Close(); err != nil {
		return err
	}
	if err := a.gzw.Close(); err != nil {
		return err
	}
	return a.file.Close()
}

// Extract extracts the tar archive at src to dest.
func ExtractTar(src, dest string) error {
	ar, err := os.Open(src)
	if err != nil {
		return err
	}
	defer ar.Close()

	gzr, err := gzip.NewReader(ar)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)
	for {
		header, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if header.Typeflag == tar.TypeReg {
			mode := header.FileInfo().Mode()
			err := extractFile(header.Name, mode, tr, dest)
			if err != nil {
				return fmt.Errorf("extract %s: %v", header.Name, err)
			}
		}
	}
}
