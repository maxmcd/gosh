package commands

import (
	"os"
	"strings"
)

// Cd changes the shell's directory
func Cd(argv []string) error {
	switch {
	case len(argv) == 1:
		return os.Chdir(os.Getenv("HOME"))
	case argv[1][0:1] == "/":
		return os.Chdir(argv[1])
	case argv[1][0:1] == "~":
		return os.Chdir(os.Getenv("HOME") + strings.Join(argv[1:], ""))
	default:
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		return os.Chdir(wd + "/" + argv[1])
	}
}
