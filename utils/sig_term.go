package utils

import (
	"os"
	"os/signal"
	"syscall"
)

type SigCallback func(sig *os.Signal)

func OnSigTerm(callback SigCallback) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for sig := range c {
			callback(&sig)
		}
	}()
}
