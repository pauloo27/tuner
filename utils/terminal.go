package utils

import (
	"os"
	"syscall"
	"unsafe"
)

type TerminalSize struct {
	Row uint16
	Col uint16
}

func GetTerminalSize() *TerminalSize {
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
