package x_testing

import (
	"fmt"
)

func FormatTest(funcName string, got, expected interface{}) string {
	return fmt.Sprintf("%s failed. Got %v, expected %v", funcName, got, expected)
}
