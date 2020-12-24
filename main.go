package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/nikoksr/nsh/internal/command"
	"github.com/nikoksr/nsh/internal/history"
)

var (
	promptSymbol    = ">"
	homeSymbol      = "~"
	browsingHistory = false

	commandHistory *history.CommandHistory
)

func setup() {
	commandHistory = history.NewCommandHistory()
}

func execInput(input string) error {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")

	// Split the input to separate the command and the arguments.
	args := strings.Split(input, " ")

	if len(args) < 1 {
		return nil
	}

	// Set together command
	cmd := command.New(args[0], args[1:])

	// Store input in history
	go commandHistory.Append(*cmd)

	return cmd.Execute(os.Stdout, os.Stderr)
}

func prettifyPath(path string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return path, err
	}

	if homeDir != "" {
		if path == homeDir {
			path = homeSymbol
		} else {
			path = strings.Replace(path, homeDir, homeSymbol, 1)
		}
	}

	return path, nil
}

func printPrompt() {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		cwd = "N/A"
		_, _ = fmt.Fprintln(os.Stderr, err)
	}

	if runtime.GOOS != "windows" {
		cwd, err = prettifyPath(cwd)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}

	fmt.Printf("\n%s\n", cwd)
	fmt.Print(promptSymbol + " ")
}

func mainLoop() {
	reader := bufio.NewReader(os.Stdin)

	for {
		printPrompt()

		// Read the keyboard input.
		input, err := reader.ReadString('\n')
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}

		if input == "^[[A" {
			browsingHistory = true
		}

		// Handle the execution of the input.
		err = execInput(input)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}
}

func main() {
	setup()
	mainLoop()
}
