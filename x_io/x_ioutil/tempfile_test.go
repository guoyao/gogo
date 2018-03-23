package x_ioutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/guoyao/gogo/x_testing"
)

func TestTempFile(t *testing.T) {
	funcName := "TempFile"
	content := "hello"
	f, err := TempFile("", "", []byte(content))
	if err != nil {
		t.Error(x_testing.Error(funcName, err, nil))
	} else {
		defer func() {
			f.Close()
			if f != nil {
				os.Remove(f.Name())
			}
		}()

		byteArray, err := ioutil.ReadAll(f)
		if err != nil {
			t.Error(x_testing.Error(funcName, err, nil))
		} else if string(byteArray) != content {
			t.Error(x_testing.Error(funcName, string(byteArray), content))
		}
	}

	pwd, err := os.Getwd()
	if err != nil {
		t.Error(x_testing.Error(funcName, err, nil))
	} else {
		content = "world"
		dir := filepath.Join(pwd, "guoyao")
		f2, err := TempFile(dir, "temp", []byte(content))
		if err != nil {
			t.Error(x_testing.Error(funcName, err, nil))
		} else {
			defer func() {
				f2.Close()
				if f2 != nil {
					os.RemoveAll(dir)
				}
			}()

			byteArray, err := ioutil.ReadAll(f2)
			if err != nil {
				t.Error(x_testing.Error(funcName, err, nil))
			} else if string(byteArray) != content {
				t.Error(x_testing.Error(funcName, string(byteArray), content))
			}
		}
	}
}

func TestTempFileWithSize(t *testing.T) {
	funcName := "TempFileWithSize"
	var size int64 = 1024
	f, err := TempFileWithSize(size)
	if err != nil {
		t.Error(x_testing.Error(funcName, err, nil))
	} else {
		defer func() {
			f.Close()
			if f != nil {
				os.Remove(f.Name())
			}
		}()

		stat, err := f.Stat()
		if err != nil {
			t.Error(x_testing.Error(funcName, err, nil))
		} else if stat.Size() != size {
			t.Error(x_testing.Error(funcName, stat.Size(), size))
		}
	}
}
