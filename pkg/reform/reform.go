// Package reform controls the logic for reformatting text into hash structures
package reform

import (
	"bufio"
	"fmt"
	"html"
	"net/url"
)

// EncodeInput URL and HTML encode standard input and prints new instances
//
// Args:
//
//	stdIn (*bufio.Scanner): Standard in as a buffer
//
// Returns:
//
//	None
func EncodeInput(stdIn *bufio.Scanner) {
	for stdIn.Scan() {
		urlEncoded, htmlEncoded, escapeEncoded := EncodeString(stdIn.Text())

		if urlEncoded != "" {
			fmt.Println(urlEncoded)
		}

		if htmlEncoded != "" {
			fmt.Println(htmlEncoded)
		}

		if escapeEncoded != "" {
			fmt.Println(escapeEncoded)
		}
	}
}

// EncodeString is used to URL and HTML encode a string where possible
//
// # Only returns if the output is different than the input string
//
// Args:
//
//	s (string): Input string
//
// Returns:
//
//	urlEncoded (string): Input string URL encoded
//	htmlEncoded (string): Input string HTML encoded
//	escapedEncoded (string): Input string ASCII escaped encoded
func EncodeString(s string) (string, string, string) {
	urlEncoded := url.QueryEscape(s)
	htmlEncoded := html.EscapeString(s)
	escapedEncoded := AsciiEscapeUnicode(s)

	if urlEncoded == s {
		urlEncoded = ""
	}

	if htmlEncoded == s {
		htmlEncoded = ""
	}

	if escapedEncoded == s {
		escapedEncoded = ""
	}

	return urlEncoded, htmlEncoded, escapedEncoded
}

// AsciiEscapeUnicode will convert a string into an ASCII escaped format
//
// Args:
//
//	str (string): String to escape
//
// Returns:
//
//	escapedRunes (string): Converted runes in string format
func AsciiEscapeUnicode(str string) string {
	runes := []rune(str)
	escapedRunes := make([]rune, 0, len(runes))

	for _, r := range runes {
		if r > 127 {
			// The rune is non-ASCII
			escapedRune := []rune(fmt.Sprintf("\\u%04x", r))
			escapedRunes = append(escapedRunes, escapedRune...)
		} else {
			escapedRunes = append(escapedRunes, r)
		}
	}
	return string(escapedRunes)
}
