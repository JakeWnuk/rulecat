// Package rule contains the CLI logic for rulecat
package rule

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jakewnuk/rulecat/pkg/utils"
)

// AppendRules will turn stdin to append rules
//
// Args:
//
//	stdIn (*bufio.Scanner): Standard input as a buffer
//	mode (string): Mode function to use to modify operation
//
// Returns:
//
//	None
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
//
// Args:
//
//	stdIn (*bufio.Scanner): Standard input as a buffer
//	mode (string): Mode function to use to modify operation
//
// Returns:
//
//	None
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
//
// Args:
//
//	stdIn (*bufio.Scanner): Standard input as a buffer
//	index (string): Integer of where to start the operation
//
// Returns:
//
//	None
func InsertRules(stdIn *bufio.Scanner, index string) {
	i, err := strconv.Atoi(index)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
	for stdIn.Scan() {
		rule := utils.CharToIteratingRule(stdIn.Text(), "i", i)
		fmt.Println(rule)
	}
}

// OverwriteRules will turn stdin to overwrite rules starting at an index
//
// Args:
//
//	stdIn (*bufio.Scanner): Standard input as a buffer
//	index (string): Integer of where to start the operation
//
// Returns:
//
//	None
func OverwriteRules(stdIn *bufio.Scanner, index string) {
	i, err := strconv.Atoi(index)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
	for stdIn.Scan() {
		rule := utils.CharToIteratingRule(stdIn.Text(), "o", i)
		fmt.Println(rule)
	}
}

// ToggleRules will turn stdin to toggle rules starting at an index
//
// Args:
//
//	stdIn (*bufio.Scanner): Standard input as a buffer
//	index (string): Integer of where to start the operation
//
// Returns:
//
//	None
func ToggleRules(stdIn *bufio.Scanner, index string) {
	i, err := strconv.Atoi(index)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
	for stdIn.Scan() {
		rule := utils.StringToToggle(stdIn.Text(), "T", i)
		if rule != "" {
			fmt.Println(rule)
		}
	}
}

// BlankLines will print a blank line for each item in stdin for -a9
//
// Args:
//
//	stdIn (*bufio.Scanner): Standard input as a buffer
//
// Returns:
//
//	None
func BlankLines(stdIn *bufio.Scanner) {
	for stdIn.Scan() {
		fmt.Println("")
	}
}

// CartesianRules will create the Caresian product of stdin and the input file
//
// # Standard input will be placed before file content
//
// Args:
//
//	stdIn (*bufio.Scanner): Standard input as a buffer
//	file ([]byte): Lines of a file that are used in the operation
//
// Returns:
//
//	None
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
//
// Args:
//
//	stdIn (*bufio.Scanner): Standard input as a buffer
//	rule (string): String that is used in the operation
//
// Returns:
//
//	None
func CharsToRules(stdIn *bufio.Scanner, rule string) {
	for stdIn.Scan() {
		rule := utils.CharToRule(stdIn.Text(), rule)
		utils.PrintCharacterRuleOutput(rule)
	}

}
