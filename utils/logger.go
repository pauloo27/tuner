package utils

import (
	"log"
	"os"
	"syscall"
	"unsafe"
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

type TerminalSize struct {
	Row uint16
	Col uint16
}

func GetTerminalWidth() *TerminalSize {
	ws := &TerminalSize{}
	retCode, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if retCode != 0 {
		HandleError(err, "Cannot get terminal size")
		os.Exit(0)
	}

	return ws
}

func HandleError(err error, message string) {
	if err != nil {
		log.Fatal(ColorRed, message, "\n", err)
	}
}
