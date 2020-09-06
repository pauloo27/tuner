package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
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
	if len(validValues) == 0 {
		fmt.Printf("%s » %s%s: ", ColorBlue, message, ColorReset)
	} else {
		fmt.Printf("%s » %s %v%s: ", ColorBlue, message, validValues, ColorReset)
	}

	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
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
	go func() {
		i := 0
		for !done {
			if clearScreen {
				MoveCursorTo(1, 1)
				ClearScreen()
			} else {
				EditLastLine()
			}

			fmt.Printf("%s%s%s %s\n", ColorBlue, brailleChars[i], ColorReset, message)
			i++

			if i >= len(brailleChars) {
				i = 0
			}
			time.Sleep(stepTime)
		}
		fmt.Printf("%s%s%s %s\n", ColorGreen, brailleFull, ColorReset, message)
		c <- true
	}()

	<-c
	done = true
}
