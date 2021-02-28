package builtin

import (
	"github.com/maxmcd/gosh/builtin/commands"
)

// Check verifies if a command is a builtin command
func Check(argv []string) func([]string) error {
	switch argv[0] {
	case "cd":
		return commands.Cd
	case "exit":
		return commands.Exit
	case "ls":
		return commands.Ls
	case "touch":
		return commands.Touch
	case "pwd":
		return commands.Pwd
	case "networktest":
		return commands.Networktest
	default:
		return nil
	}
}
