package utils

import "fmt"

func Fmt(format string, v ...interface{}) string {
	return fmt.Sprintf(format, v...)
}
