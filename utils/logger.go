package utils

import "log"

const ColorBold = "\033[1m"
const ColorReset = "\033[0m"
const ColorRed = "\033[31m"
const ColorGreen = "\033[32m"
const ColorYellow = "\033[33m"
const ColorBlue = "\033[34m"
const ColorWhite = "\033[39m"

func HandleError(err error, message string) {
	if err != nil {
		log.Fatal(ColorRed, message, "\n", err)
	}
}
