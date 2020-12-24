package command

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type Command struct {
	Name       string
	Args       []string
	IsBuiltIn  bool
	MinNumArgs int
	MaxNumArgs int
	runFunc
}

func isCommandBuiltIn(name string) (*Command, bool) {
	for _, cmd := range builtInCommands {
		if cmd.Name == name {
			return &cmd, true
		}
	}

	return nil, false
}

func New(name string, args []string) *Command {
	cmd := &Command{
		Name: name,
		Args: args,
	}

	builtInCmd, isBuiltIn := isCommandBuiltIn(name)
	if isBuiltIn && builtInCmd != nil {
		cmd = builtInCmd
	}

	return cmd
}

func (c Command) Execute(stdout, stderr io.Writer) error {
	var err error

	if c.IsBuiltIn {
		err = c.executeBuiltIn()
	} else {
		err = c.executeProgram(stdout, stderr)
	}

	return err
}

func (c Command) executeProgram(stdout, stderr io.Writer) error {
	// Pass the program and the arguments separately.
	cmd := exec.Command(c.Name, c.Args...)

	// Set the correct output device.
	cmd.Stdout = stderr
	cmd.Stderr = stdout

	// Execute the command and return the error.
	return cmd.Run()
}

func (c Command) executeBuiltIn() error {
	return c.runFunc(c.Args...)
}

func (c Command) ToString() string {
	return fmt.Sprintf("%s %s", c.Name, strings.Join(c.Args, " "))
}
