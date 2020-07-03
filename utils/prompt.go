package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)
var brailleChars = [8]string{"⡿", "⣟", "⣯", "⣷", "⣾", "⣽", "⣻", "⢿"}
var brailleFull = "⣿"

func AskFor(message string, validValues ...string) string {
	if len(validValues) == 0 {
		fmt.Printf("%s » %s%s: ", ColorBlue, message, ColorReset)
	} else {
		fmt.Printf("%s » %s %v%s: ", ColorBlue, message, validValues, ColorReset)
	}

	line, err := reader.ReadString('\n')
	HandleError(err, "Cannot read user input")

	response := strings.TrimSuffix(line, "\n")
	if len(validValues) == 0 {
		return response
	}

	for _, value := range validValues {
		if strings.EqualFold(value, response) {
			return value
		}
	}
	HandleError(fmt.Errorf("Invalid response. Valid responses are %v.", validValues), "Invalid response")
	return ""
}

func EditLastLine() {
	fmt.Printf("\x1b[1F")
}

func PrintWithLoadIcon(message string, c chan bool) {
	done := false
	go func() {
		i := 0
		for !done {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("%s%s%s %s\n", ColorBlue, brailleChars[i], ColorReset, message)
			EditLastLine()
			i++

			if i >= len(brailleChars) {
				i = 0
			}
		}
		fmt.Printf("%s%s%s %s\n", ColorGreen, brailleFull, ColorReset, message)
	}()

	<-c
	done = true
}
