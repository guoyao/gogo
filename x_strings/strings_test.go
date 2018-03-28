package x_strings

import (
	"testing"

	"github.com/guoyao/gogo/x_testing"
)

func TestIsBlank(t *testing.T) {
	funcName := "IsBlank"
	expected, result := true, IsBlank("")
	if result != expected {
		t.Error(x_testing.Error(funcName, result, expected))
	}

	expected, result = true, IsBlank("  ")
	if result != expected {
		t.Error(x_testing.Error(funcName, result, expected))
	}

	expected, result = false, IsBlank(" 1 ")
	if result != expected {
		t.Error(x_testing.Error(funcName, result, expected))
	}
}

func TestLen(t *testing.T) {
	funcName := "Len"
	expected, result := 12, Len("hello golang")
	if result != expected {
		t.Error(x_testing.Error(funcName, result, expected))
	}

	expected, result = 2, Len("你好")
	if result != expected {
		t.Error(x_testing.Error(funcName, result, expected))
	}

	expected, result = 9, Len("你好，golang")
	if result != expected {
		t.Error(x_testing.Error(funcName, result, expected))
	}
}

func TestAt(t *testing.T) {
	funcName := "At"
	expected, result := "g", At("hello golang", 6)
	if result != expected {
		t.Error(x_testing.Error(funcName, result, expected))
	}

	expected, result = "好", At("你好", 1)
	if result != expected {
		t.Error(x_testing.Error(funcName, result, expected))
	}

	expected, result = "好", At("hi, 你abc好，golang", 8)
	if result != expected {
		t.Error(x_testing.Error(funcName, result, expected))
	}
}

func TestIndex(t *testing.T) {
	funcName := "Index"
	expected, result := 2, Index("golang", "lang")
	if result != expected {
		t.Error(x_testing.Error(funcName, result, expected))
	}

	expected, result = 3, Index("Go你好", "好")
	if result != expected {
		t.Error(x_testing.Error(funcName, result, expected))
	}

	expected, result = 0, Index("你好", "你")
	if result != expected {
		t.Error(x_testing.Error(funcName, result, expected))
	}

	expected, result = -1, Index("Go你好", "好啊")
	if result != expected {
		t.Error(x_testing.Error(funcName, result, expected))
	}
}
