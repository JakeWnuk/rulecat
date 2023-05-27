// Package that contains the primary logic for rulecat and the CLI
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/jakewnuk/rulecat/pkg/utils"
)

func main() {
	if len(os.Args) <= 1 {
		printUsage()
		os.Exit(0)
	}

	stdIn := bufio.NewScanner(os.Stdin)

	_, err := os.Stat(os.Args[1])
	if err == nil {
		file, err := ioutil.ReadFile(os.Args[1])
		utils.CheckError(err)
		cartesianRules(stdIn, file)
		os.Exit(0)
	}

	switch os.Args[1] {
	case "append":
		// if no mode use default
		if len(os.Args) == 2 {
			os.Args = append(os.Args, "default")
		}
		appendRules(stdIn, os.Args[2])
	case "prepend":
		// if no mode use default
		if len(os.Args) == 2 {
			os.Args = append(os.Args, "default")
		}
		prependRules(stdIn, os.Args[2])
	case "insert":
		if len(os.Args) == 2 {
			os.Args = append(os.Args, "0")
		}
		insertRules(stdIn, os.Args[2])
	case "overwrite":
		if len(os.Args) == 2 {
			os.Args = append(os.Args, "0")
		}
		overwriteRules(stdIn, os.Args[2])
	case "toggle":
		if len(os.Args) == 2 {
			os.Args = append(os.Args, "0")
		}
		toggleRules(stdIn, os.Args[2])
	case "blank":
		blankLines(stdIn)
	case "chars":
		charsToRules(stdIn, os.Args[2])
	default:
		printUsage()
		os.Exit(0)
	}
}

// printUsage prints usage information for the program
func printUsage() {
	fmt.Println("OPTIONS: append prepend blank <RULE-FILE> chars insert overwrite toggle")
	fmt.Println("EXAMPLE: stdin | rulecat append")
	fmt.Println("EXAMPLE: stdin | rulecat prepend")
	fmt.Println("EXAMPLE: stdin | rulecat append remove")
	fmt.Println("EXAMPLE: stdin | rulecat prepend remove")
	fmt.Println("EXAMPLE: stdin | rulecat append shift")
	fmt.Println("EXAMPLE: stdin | rulecat prepend shift")
	fmt.Println("EXAMPLE: stdin | rulecat blank")
	fmt.Println("EXAMPLE: stdin | rulecat <RULE-FILE>")
	fmt.Println("EXAMPLE: stdin | rulecat chars <RULE>")
	fmt.Println("EXAMPLE: stdin | rulecat insert <START-INDEX>")
	fmt.Println("EXAMPLE: stdin | rulecat overwrite <START-INDEX>")
	fmt.Println("EXAMPLE: stdin | rulecat toggle <START-INDEX>")
}

// appendRules will turn stdin to append rules
func appendRules(stdIn *bufio.Scanner, mode string) {
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

// prependRules will turn stdin to prepend rules
func prependRules(stdIn *bufio.Scanner, mode string) {
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

// insertRules will turn stdin to insert rules starting at an index
func insertRules(stdIn *bufio.Scanner, index string) {
	i, err := strconv.Atoi(index)
	utils.CheckError(err)
	for stdIn.Scan() {
		rule := utils.CharToIteratingRule(stdIn.Text(), "i", i)
		fmt.Println(rule)
	}
}

// overwriteRules will turn stdin to overwrite rules starting at an index
func overwriteRules(stdIn *bufio.Scanner, index string) {
	i, err := strconv.Atoi(index)
	utils.CheckError(err)
	for stdIn.Scan() {
		rule := utils.CharToIteratingRule(stdIn.Text(), "o", i)
		fmt.Println(rule)
	}
}

// toggleRules will turn stdin to toggle rules starting at an index
func toggleRules(stdIn *bufio.Scanner, index string) {
	i, err := strconv.Atoi(index)
	utils.CheckError(err)
	for stdIn.Scan() {
		rule := utils.StringToToggle(stdIn.Text(), "T", i)
		if rule != "" {
			fmt.Println(rule)
		}
	}
}

// blankLines will print a blank line for each item in stdin for -a9
func blankLines(stdIn *bufio.Scanner) {
	for stdIn.Scan() {
		fmt.Println("")
	}
}

// cartesianRules will create the Caresian product of stdin and the input file
// NOTE: stdin will be placed before file content
func cartesianRules(stdIn *bufio.Scanner, file []byte) {
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

// charsToRules will insert a custom rule before each character
func charsToRules(stdIn *bufio.Scanner, rule string) {
	for stdIn.Scan() {
		rule := utils.CharToRule(stdIn.Text(), rule)
		utils.PrintCharacterRuleOutput(rule)
	}

}
