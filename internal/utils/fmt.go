package utils

import (
	"fmt"

	"github.com/rivo/tview"
)

func Fmt(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func FmtEscaping(format string, args ...interface{}) string {
	var escapedArgs []interface{}

	for _, arg := range args {
		switch a := arg.(type) {
		case string:
			escapedArgs = append(escapedArgs, tview.Escape(a))
		default:
			escapedArgs = append(escapedArgs, a)
		}
	}

	return fmt.Sprintf(format, escapedArgs...)
}
