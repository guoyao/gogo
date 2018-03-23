package x_ioutil

import (
	"io/ioutil"
	"os"

	"github.com/guoyao/gogo/x_os"
	"github.com/guoyao/gogo/x_strings"
)

// TempFile generates a temp file with separated content.
func TempFile(dir, prefix string, content []byte) (*os.File, error) {
	if !x_strings.IsBlank(dir) && !x_os.IsExist(dir) {
		err := os.MkdirAll(dir, 0700)
		if err != nil {
			return nil, err
		}
	}

	f, err := ioutil.TempFile(dir, prefix)
	if err != nil {
		return f, err
	}

	if content != nil {
		if _, err := f.Write(content); err != nil {
			return f, err
		}
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		return f, err
	}

	return f, nil
}

// TempFileWithSize generates a temp file with specified size.
func TempFileWithSize(fileSize int64) (*os.File, error) {
	f, err := TempFile("", "", nil)
	if err != nil {
		return f, err
	}

	if err = f.Truncate(fileSize); err != nil {
		return f, err
	}

	return f, nil
}
