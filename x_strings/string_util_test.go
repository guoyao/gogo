package x_strings

import (
	"testing"

	"github.com/guoyao/gogo/x_testing"
)

func TestIndexOf(t *testing.T) {
	funcName, expected, result := "IndexOf", 2, IndexOf("golang", "lang")
	if result != expected {
		t.Error(x_testing.FormatTest(funcName, result, expected))
	}

	expected, result = 3, IndexOf("Go你好", "好")
	if result != expected {
		t.Error(x_testing.FormatTest(funcName, result, expected))
	}

	expected, result = 0, IndexOf("你好", "你")
	if result != expected {
		t.Error(x_testing.FormatTest(funcName, result, expected))
	}

	expected, result = -1, IndexOf("Go你好", "好啊")
	if result != expected {
		t.Error(x_testing.FormatTest(funcName, result, expected))
	}
}
