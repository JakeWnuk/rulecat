// Package that contains the primary logic for rulecat and the CLI
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jakewnuk/rulecat/pkg/reform"
	"github.com/jakewnuk/rulecat/pkg/rule"
)

var version = "0.0.1"

func main() {

	if len(os.Args) <= 1 {
		printUsage()
		os.Exit(0)
	}

	stdIn := bufio.NewScanner(os.Stdin)

	_, err := os.Stat(os.Args[1])
	if err == nil {
		file, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			os.Exit(1)
		}
		rule.CartesianRules(stdIn, file)
		os.Exit(0)
	}

	switch os.Args[1] {
	case "append":
		// if no mode use default
		if len(os.Args) == 2 {
			os.Args = append(os.Args, "default")
		}
		rule.AppendRules(stdIn, os.Args[2])
	case "prepend":
		// if no mode use default
		if len(os.Args) == 2 {
			os.Args = append(os.Args, "default")
		}
		rule.PrependRules(stdIn, os.Args[2])
	case "insert":
		if len(os.Args) == 2 {
			os.Args = append(os.Args, "0")
		}
		rule.InsertRules(stdIn, os.Args[2])
	case "overwrite":
		if len(os.Args) == 2 {
			os.Args = append(os.Args, "0")
		}
		rule.OverwriteRules(stdIn, os.Args[2])
	case "toggle":
		if len(os.Args) == 2 {
			os.Args = append(os.Args, "0")
		}
		rule.ToggleRules(stdIn, os.Args[2])
	case "blank":
		rule.BlankLines(stdIn)
	case "chars":
		rule.CharsToRules(stdIn, os.Args[2])
	case "encode":
		reform.EncodeInput(stdIn)
	case "combo":
		if len(os.Args) < 4 {
			fmt.Println("ERROR: Must provide 2 arguments for combo mode (toggle, prepend, append, insert)")
			os.Exit(0)
		}
		rule.ComboRules(stdIn, os.Args[2], os.Args[3])

	default:
		printUsage()
		os.Exit(0)
	}
}

// printUsage prints usage information for the program
func printUsage() {
	fmt.Println(fmt.Sprintf("\nModes for rulecat (version %s):", version))
	fmt.Println("\n  append\tCreates append rules from text")
	fmt.Println("\t\tExample: stdin | rulecat append")
	fmt.Println("\t\tExample: stdin | rulecat append remove")
	fmt.Println("\t\tExample: stdin | rulecat append shift")
	fmt.Println("\n  prepend\tCreates prepend rules from text")
	fmt.Println("\t\tExample: stdin | rulecat prepend")
	fmt.Println("\t\tExample: stdin | rulecat prepend remove")
	fmt.Println("\t\tExample: stdin | rulecat prepend shift")
	fmt.Println("\n  blank\t\tCreates blank lines from text")
	fmt.Println("\t\tExample: stdin | rulecat blank")
	fmt.Println("\n  [RULE-FILE]\tCreate cartesian product of a file and text")
	fmt.Println("\t\tExample: stdin | rulecat [FILE]")
	fmt.Println("\n  chars\t\tCreates custom rules per character from text")
	fmt.Println("\t\tExample: stdin | rulecat chars [RULE]")
	fmt.Println("\n  insert\tCreates insert rules from from text")
	fmt.Println("\t\tExample: stdin | rulecat insert [START-INDEX]")
	fmt.Println("\n  overwrite\tCreates overwrite rules from from text")
	fmt.Println("\t\tExample: stdin | rulecat overwrite [START-INDEX]")
	fmt.Println("\n  toggle\tCreates toggle rules from from text")
	fmt.Println("\t\tExample: stdin | rulecat toggle [START-INDEX]")
	fmt.Println("\n  encode\tURL, HTML, and Unicode escape encodes input and prints new output")
	fmt.Println("\t\tExample: stdin | rulecat encode")
	fmt.Println("\n  combo\t\tCombines multiple modes into one rule per line (toggle, prepend, append, insert)")
	fmt.Println("\t\tExample: stdin | rulecat combo [MODE-A] [MODE-B]")
}
