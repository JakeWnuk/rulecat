// Package cmd contains the CLI logic for rulecat
package cmd

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/jakewnuk/rulecat/pkg/utils"
)

// AppendRules will turn stdin to append rules
func AppendRules(stdIn *bufio.Scanner, mode string) {
	switch mode {
	// remove will remove characters then append
	case "remove":
		for stdIn.Scan() {
			rule := utils.CharToRule(stdIn.Text(), "$")
			remove := utils.LenToRule(stdIn.Text(), "]")
			utils.PrintCharacterRuleOutput(remove, rule)
		}
	// shift will shift characters back to front then append
	case "shift":
		for stdIn.Scan() {
			rule := utils.CharToRule(stdIn.Text(), "$")
			shift := utils.LenToRule(stdIn.Text(), "}")
			utils.PrintCharacterRuleOutput(shift, rule)
		}
	default:
		for stdIn.Scan() {
			rule := utils.CharToRule(stdIn.Text(), "$")
			utils.PrintCharacterRuleOutput(rule)
		}
	}

}

// PrependRules will turn stdin to prepend rules
func PrependRules(stdIn *bufio.Scanner, mode string) {
	switch mode {
	// remove will remove characters then prepend
	case "remove":
		for stdIn.Scan() {
			rule := utils.CharToRule(utils.ReverseString(stdIn.Text()), "^")
			remove := utils.LenToRule(stdIn.Text(), "[")
			utils.PrintCharacterRuleOutput(remove, rule)
		}
	// shift will shift characters front to back then prepend
	case "shift":
		for stdIn.Scan() {
			rule := utils.CharToRule(utils.ReverseString(stdIn.Text()), "^")
			shift := utils.LenToRule(stdIn.Text(), "{")
			utils.PrintCharacterRuleOutput(shift, rule)
		}
	default:
		for stdIn.Scan() {
			rule := utils.CharToRule(utils.ReverseString(stdIn.Text()), "^")
			utils.PrintCharacterRuleOutput(rule)
		}
	}

}

// InsertRules will turn stdin to insert rules starting at an index
func InsertRules(stdIn *bufio.Scanner, index string) {
	i, err := strconv.Atoi(index)
	utils.CheckError(err)
	for stdIn.Scan() {
		rule := utils.CharToIteratingRule(stdIn.Text(), "i", i)
		fmt.Println(rule)
	}
}

// OverwriteRules will turn stdin to overwrite rules starting at an index
func OverwriteRules(stdIn *bufio.Scanner, index string) {
	i, err := strconv.Atoi(index)
	utils.CheckError(err)
	for stdIn.Scan() {
		rule := utils.CharToIteratingRule(stdIn.Text(), "o", i)
		fmt.Println(rule)
	}
}

// ToggleRules will turn stdin to toggle rules starting at an index
func ToggleRules(stdIn *bufio.Scanner, index string) {
	i, err := strconv.Atoi(index)
	utils.CheckError(err)
	for stdIn.Scan() {
		rule := utils.StringToToggle(stdIn.Text(), "T", i)
		if rule != "" {
			fmt.Println(rule)
		}
	}
}

// BlankLines will print a blank line for each item in stdin for -a9
func BlankLines(stdIn *bufio.Scanner) {
	for stdIn.Scan() {
		fmt.Println("")
	}
}

// CartesianRules will create the Caresian product of stdin and the input file
// NOTE: stdin will be placed before file content
func CartesianRules(stdIn *bufio.Scanner, file []byte) {
	fileLines := strings.Split(string(file), "\n")
	for stdIn.Scan() {
		input := stdIn.Text()
		for _, line := range fileLines {
			if line != "" {
				fmt.Printf("%s %s\n", input, line)
			}
		}
	}
}

// CharsToRules will insert a custom rule before each character
func CharsToRules(stdIn *bufio.Scanner, rule string) {
	for stdIn.Scan() {
		rule := utils.CharToRule(stdIn.Text(), rule)
		utils.PrintCharacterRuleOutput(rule)
	}

}
