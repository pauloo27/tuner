package core

import (
	"fmt"

	"github.com/rivo/tview"
)

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
