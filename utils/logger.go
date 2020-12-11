package utils

import (
	"log"
)

const (
	ColorBold   = "\033[1m"
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorWhite  = "\033[39m"
)

func HandleError(err error, message string) {
	if err != nil {
		log.Fatal(ColorRed, message, "\n", err)
	}
}
