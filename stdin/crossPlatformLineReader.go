package stdin

import (
	"bufio"
	"os"
	"strings"
)

var reader *bufio.Reader

func ReadLine() (string, error) {

	// Initializing reader on first usage
	if reader == nil {
		reader = bufio.NewReader(os.Stdin)
	}

	// Reading line and error
	text, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	// convert CRLF to LF
	text = strings.Replace(text, "\r", "", -1)

	// Remove the last character LF
	text = text[0 : len(text)-1]

	// Trim space around
	text = strings.TrimSpace(text)

	return text, nil
}
