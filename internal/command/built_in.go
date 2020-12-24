package command

import "os"

type runFunc = func(...string) error

var (
	builtInCommands = []Command{
		cd,
		exit,
	}

	cd = Command{
		Name:       "cd",
		Args:       []string{},
		IsBuiltIn:  true,
		MinNumArgs: 0,
		MaxNumArgs: 1,
		runFunc:    runCd,
	}

	exit = Command{
		Name:       "exit",
		Args:       []string{},
		IsBuiltIn:  true,
		MinNumArgs: 0,
		MaxNumArgs: 0,
		runFunc:    runExit,
	}
)

func runCd(paths ...string) error {
	path := "~"
	if len(paths) > 0 && len(paths[0]) > 0 {
		path = paths[0]
	}

	return os.Chdir(path)
}

func runExit(args ...string) error {
	os.Exit(0)

	return nil
}
