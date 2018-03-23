package x_os

import (
	"testing"

	"github.com/guoyao/gogo/x_testing"
)

func TestIsExist(t *testing.T) {
	funcName := "IsExist"
	expected, result := true, IsExist("os_test.go")
	if result != expected {
		t.Error(x_testing.Error(funcName, result, expected))
	}

	expected, result = false, IsExist("not_exist_file.go")
	if result != expected {
		t.Error(x_testing.Error(funcName, result, expected))
	}

	expected, result = false, IsExist("")
	if result != expected {
		t.Error(x_testing.Error(funcName, result, expected))
	}
}
