package stdInput

import (
	"bufio"
	"os"
	"strings"
)

var reader *bufio.Reader

func StartReader() {
	reader = bufio.NewReader(os.Stdin)
}

func ReadLine() (string, error) {
	// Reading line and error (line to 'text' and error to '_')
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
