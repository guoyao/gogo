package x_testing

import (
	"fmt"
)

func Error(name, got, expected interface{}) string {
	return fmt.Sprintf("%v failed. Got %v, expected %v", name, got, expected)
}
