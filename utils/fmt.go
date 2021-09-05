package utils

import "fmt"

func Fmt(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
