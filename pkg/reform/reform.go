// Package reform controls the logic for reformatting text into hash structures
package reform

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"html"
	"net/url"
	"strconv"
	"strings"
)

// hashOperation is a struct used to hold the parsed expression data
//
// The struct has the following fields:
//
//	algorithm: The algorithm to use for hashing
//	times: The number of times it should be used
type hashOperation struct {
	algorithm string
	times     int
}

// hashString transforms an input string by a hashOperation struct
//
// Args:
//
//	op (hashOperation): Struct containing the transformations to make
//	input (string): Input string to transform
//
// Returns:
//
//	input (string): Transformed string
func hashString(op hashOperation, input string) string {
	for i := 0; i < op.times; i++ {
		var h hash.Hash
		switch op.algorithm {
		case "md5":
			h = md5.New()
		case "sha1":
			h = sha1.New()
		case "sha256":
			h = sha256.New()
		case "sha512":
			h = sha512.New()
		default:
			panic("unsupported algorithm: " + op.algorithm)
		}
		h.Write([]byte(input))
		input = hex.EncodeToString(h.Sum(nil))
	}
	return input
}

// parseHashExpression is used to parse user provided expression into
// operations
//
// General format is: 2xmd5(p) or 100xsha256(md5(p))
//
// Args:
//
//	expr (string): Expression string
//
// Returns:
//
//	operations ([]hashOperation): Parsed operations to use
func parseHashExpression(expr string) []hashOperation {
	var operations []hashOperation
	for strings.Contains(expr, "(") && strings.Contains(expr, ")") {
		start := strings.Index(expr, "(")
		end := strings.LastIndex(expr, ")")
		algorithm := expr[:start]
		if strings.Contains(algorithm, "x") {
			parts := strings.Split(algorithm, "x")
			algorithm = parts[1]
			times, err := strconv.Atoi(parts[0])
			if err != nil {
				panic("invalid times: " + parts[0])
			}
			operations = append(operations, hashOperation{algorithm, times})
		} else {
			operations = append(operations, hashOperation{algorithm, 1})
		}
		expr = expr[start+1 : end]
	}
	return operations
}

// RehashByExpression is used to rehash a plaintext by a provided expression
//
// Args:
//
//	stdIn (*bufio.Scanner): Standard in as a buffer
//	expression (string): Input expression to use
//
// Returns:
//
//	None
func RehashByExpression(stdIn *bufio.Scanner, expression string) {
	for stdIn.Scan() {
		operations := parseHashExpression(expression)
		input := string(stdIn.Text())
		for i := len(operations) - 1; i >= 0; i-- {
			input = hashString(operations[i], input)
		}
		fmt.Println(input)
	}
}

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
