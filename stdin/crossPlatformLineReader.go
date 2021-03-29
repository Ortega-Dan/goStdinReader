package stdin

import (
	"bufio"
	"os"
	"strings"
)

var reader *bufio.Reader

// Reads a line from std input without default trim
func ReadLineNoTrim() string {

	// Initializing reader on first usage
	if reader == nil {
		reader = bufio.NewReader(os.Stdin)
	}

	// Reading line and error
	text, err := reader.ReadString('\n')

	// in case of error, just panic
	if err != nil {
		panic(err)
	}

	// convert CRLF to LF
	text = strings.Replace(text, "\r", "", -1)

	// Remove the last character LF
	text = text[0 : len(text)-1]

	return text
}

// Reads a line from std input with default trim
func ReadLine() string {

	text := ReadLineNoTrim()

	// Trim space around
	text = strings.TrimSpace(text)

	return text
}
