package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	abspath string
}

func newZipEntry(path string) *ZipEntry {
	abspath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{abspath}
}
func (z ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(z.abspath)
	if err != nil {
		panic(err)
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				panic(err)
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				panic(err)
			}
			return data, z, nil
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}
func (z ZipEntry) String() string {
	return z.abspath
}
