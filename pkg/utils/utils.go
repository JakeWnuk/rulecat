// Package utils contains functions for the main rulecat program
package utils

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// LenToRule converts a string to a rule by its length
//
// Args:
//
//	str (string): Input string to transform
//	rule (string): Rule to insert per length
//
// Returns:
//
//	(string): Transformed string
func LenToRule(str string, rule string) string {
	return strings.TrimSpace(strings.Repeat(rule+" ", len(str)))
}

// CharToRule converts a string to a rule by its characters
//
// Args:
//
//	str (string): Input string to transform
//	rule (string): Rule to insert per character
//
// Returns:
//
//	(string): Transformed string
func CharToRule(str string, rule string) string {
	return rule + strings.Join(strings.Split(str, ""), " "+rule)
}

// CharToIteratingRule converts a string to a rule by its characters but
// increments along with each character
//
// Args:
//
//	str (string): Input string to transform
//	rule (string): Rule to insert per length
//	index (int): Index to start at
//
// Returns:
//
//	(string): Transformed string
func CharToIteratingRule(str string, rule string, index int) string {
	var result strings.Builder
	for i, r := range str {
		if i+index < 10 {
			result.WriteString(fmt.Sprintf("%s%d%c ", rule, i+index, r))
		} else {
			result.WriteString(fmt.Sprintf("%s%c%c ", rule, 'A'+i+index-10, r))
		}
	}
	return strings.TrimSpace(result.String())
}

// StringToToggle converts a string to toggle rules by looking for upper chars
//
// Args:
//
//	str (string): Input string to transform
//	rule (string): Rule to insert per length
//	index (int): Index to start at
//
// Returns:
//
//	(string): Transformed string
func StringToToggle(str string, rule string, index int) string {
	var result strings.Builder
	for i, r := range str {
		if unicode.IsUpper(r) {
			if i+index < 10 {
				result.WriteString(fmt.Sprintf("%s%d ", rule, i+index))
			} else if i+index-10 < 26 {
				result.WriteString(fmt.Sprintf("%s%c ", rule, 'A'+i+index-10))
			}
		}
	}
	return strings.TrimSpace(result.String())
}

// ReverseString will return a string in reverse
//
// Args:
//
//	str (string): Input string to transform
//
// Returns:
//
//	(string): Transformed string
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// CheckASCIIString checks to see if a string only contains ascii characters
//
// Args:
//
//	str (string): Input string to check
//
// Returns:
//
//	(bool): If the string only contained ASCII characters
func CheckASCIIString(str string) bool {
	if utf8.RuneCountInString(str) != len(str) {
		return false
	}
	return true
}

// PrintCharacterRuleOutput handles printing the rules to the CLI
//
// prints for CharToRule functions
//
// Args:
//
//	strs (...string): Input strings to print
//
// Returns:
//
// None
func PrintCharacterRuleOutput(strs ...string) {
	output := ""
	for _, str := range strs {
		if CheckASCIIString(str) {
			output += str + " "
		} else {
			output += ConvertCharacterMultiByteString(str)
		}
	}

	if output != "" && len(output) <= 93 {
		fmt.Println(strings.TrimSpace(output))
	}
}

// ConvertCharacterMultiByteString converts non-ascii characters to a hashcat valid format
//
// converts for CharToRule functions
//
// Args:
//
//	str (string): Input string to transform
//
// Returns:
//
//	returnStr (string): Converted string
func ConvertCharacterMultiByteString(str string) string {
	returnStr := ""
	deletedChar := ``
	for i, r := range str {
		if r > 127 {
			if i > 0 {
				deletedChar = string(returnStr[len(returnStr)-1])
				returnStr = returnStr[:len(returnStr)-1]
			}
			byteArr := []byte(string(r))
			if deletedChar == "^" {
				for j := len(byteArr) - 1; j >= 0; j-- {
					b := byteArr[j]
					if j == 0 {
						returnStr += fmt.Sprintf("%s\\x%X", deletedChar, b)
					} else {
						returnStr += fmt.Sprintf("%s\\x%X ", deletedChar, b)
					}
				}
			} else {
				for j, b := range byteArr {
					if j == len(byteArr)-1 {
						returnStr += fmt.Sprintf("%s\\x%X", deletedChar, b)
					} else {
						returnStr += fmt.Sprintf("%s\\x%X ", deletedChar, b)
					}
				}
			}
		} else {
			returnStr += fmt.Sprintf("%c", r)
		}
	}
	return returnStr
}
