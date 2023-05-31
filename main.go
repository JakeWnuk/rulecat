// Package that contains the primary logic for rulecat and the CLI
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jakewnuk/rulecat/pkg/cmd"
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
		cmd.CartesianRules(stdIn, file)
		os.Exit(0)
	}

	switch os.Args[1] {
	case "append":
		// if no mode use default
		if len(os.Args) == 2 {
			os.Args = append(os.Args, "default")
		}
		cmd.AppendRules(stdIn, os.Args[2])
	case "prepend":
		// if no mode use default
		if len(os.Args) == 2 {
			os.Args = append(os.Args, "default")
		}
		cmd.PrependRules(stdIn, os.Args[2])
	case "insert":
		if len(os.Args) == 2 {
			os.Args = append(os.Args, "0")
		}
		cmd.InsertRules(stdIn, os.Args[2])
	case "overwrite":
		if len(os.Args) == 2 {
			os.Args = append(os.Args, "0")
		}
		cmd.OverwriteRules(stdIn, os.Args[2])
	case "toggle":
		if len(os.Args) == 2 {
			os.Args = append(os.Args, "0")
		}
		cmd.ToggleRules(stdIn, os.Args[2])
	case "blank":
		cmd.BlankLines(stdIn)
	case "chars":
		cmd.CharsToRules(stdIn, os.Args[2])
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
