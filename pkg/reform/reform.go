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
	"regexp"
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

// DehexInput is used to dehex input from $HEX[...] format
//
// Args:
//
//	stdIn (*bufio.Scanner): Standard in as a buffer
//
// Returns:
//
//	None
func DehexInput(stdIn *bufio.Scanner) {
	for stdIn.Scan() {
		if TestHexInput(stdIn.Text()) == true {
			plaintext, err := DehexPlaintext(stdIn.Text())
			if err != nil {
				panic("invalid hex input: " + plaintext)
			}
			fmt.Println(plaintext)
		}
	}
}

// DehexPlaintext decodes plaintext from $HEX[...] format
//
// Args:
//
//	s (string): The string to be dehexed
//
// Returns:
//
//	decoded (string): The decoded hex string
//	err (error): Error data
func DehexPlaintext(s string) (string, error) {
	s = strings.TrimPrefix(s, "$HEX[")
	s = strings.TrimSuffix(s, "]")
	decoded, err := hex.DecodeString(s)
	return string(decoded), err
}

// TestHexInput is used by the rehashing feature to identify plaintext in the
// $HEX[...] format
//
// Args:
//
//	s (str): The string to be evaluated
//
// Returns:
//
//	(bool): Returns true if it matches and false if it did not
func TestHexInput(s string) bool {
	var validateInput = regexp.MustCompile(`^\$HEX\[[a-zA-Z0-9]*\]$`).MatchString
	if validateInput(s) == false {
		return false
	}
	return true
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
		urlEncoded, htmlEncoded := EncodeString(stdIn.Text())

		if urlEncoded != "" {
			fmt.Println(urlEncoded)
		}

		if htmlEncoded != "" {
			fmt.Println(htmlEncoded)
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
func EncodeString(s string) (string, string) {
	urlEncoded := url.QueryEscape(s)
	htmlEncoded := html.EscapeString(s)

	if urlEncoded == s {
		urlEncoded = ""
	}

	if htmlEncoded == s {
		htmlEncoded = ""
	}

	return urlEncoded, htmlEncoded
}
