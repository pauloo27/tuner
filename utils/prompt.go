package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/term"
)

var reader = bufio.NewReader(os.Stdin)
var brailleChars = [8]string{"⢿", "⣻", "⣽", "⣾", "⣷", "⣯", "⣟", "⡿"}
var brailleFull = "⣿"

func WaitForEnter(message string) error {
	fmt.Printf("%s%s%s", ColorGreen, message, ColorReset)
	_, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	return nil
}

func AskFor(message string, validValues ...string) (string, error) {
	fd := int(os.Stdout.Fd())
	isTerminal := term.IsTerminal(fd)

	if isTerminal {
		oldState, err := term.MakeRaw(fd)
		if err != nil {
			return "", err
		}

		defer term.Restore(fd, oldState)
	}

	if len(validValues) == 0 {
		fmt.Printf("%s > %s%s: ", ColorBlue, message, ColorReset)
	} else {
		fmt.Printf("%s > %s %v%s: ", ColorBlue, message, validValues, ColorReset)
	}

	var err error
	var line string

	if isTerminal {
		terminal := term.NewTerminal(os.Stdout, "")
		line, err = terminal.ReadLine()
		if err != nil {
			return "", err
		}
	} else {
		line, err = reader.ReadString('\n')
		if err != nil {
			return "", err
		}
	}

	response := strings.TrimSuffix(line, "\n")
	if len(validValues) == 0 {
		return response, nil
	}

	for _, value := range validValues {
		if strings.EqualFold(value, response) {
			return value, nil
		}
	}
	return "", fmt.Errorf("Invalid response. Valid responses are %v.", validValues)
}

func AskForInt(message string) (int, error) {
	input, err := AskFor(message)
	if err != nil {
		return 0, err
	}

	v, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func ClearScreen() {
	MoveCursorTo(1, 1)
	ClearAfterCursor()
}

func ClearAfterCursor() {
	fmt.Printf("\033[J")
}

func MoveCursorTo(line, column int) {
	fmt.Printf("\033[%d;%df", line, column)
}

func MoveCursorUp(lineCount int) {
	fmt.Printf("\x1b[%dF", lineCount)
}

func EditLastLine() {
	MoveCursorUp(1)
}

func PrintWithLoadIcon(message string, c chan bool, stepTime time.Duration, clearScreen bool) {
	done := false

	print := func(format string, v ...interface{}) {
		if clearScreen {
			MoveCursorTo(1, 1)
			ClearScreen()
		} else {
			EditLastLine()
		}
		fmt.Printf(format, v...)
	}

	go func() {
		i := 0
		for !done {
			print("%s%s%s %s\n", ColorBlue, brailleChars[i], ColorReset, message)
			i++

			if i >= len(brailleChars) {
				i = 0
			}
			time.Sleep(stepTime)
		}
		print("%s%s%s %s\n", ColorGreen, brailleFull, ColorReset, message)
		c <- true
	}()

	<-c
	done = true
}
