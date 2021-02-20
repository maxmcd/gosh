package shell

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/maxmcd/gosh/builtin"
)

func interpret() error {
	// todo: add signal listener

	// print the prompt
	shellPrompt.Print()

	// read user input
	input, err := readInput()

	// return errors
	if err != nil {
		if err == io.EOF {
			return err
		}
		return err
	}

	input = strings.TrimRight(input, "\r\n")

	// if no input is given, skip the cycle
	if input == "" {
		return nil
	}

	// todo: add history

	// separate the input in arguments
	argv := strings.Fields(input)

	// check if the command is a builtin command
	fn := builtin.Check(argv)
	if fn != nil {
		return fn(argv)
	}
	fmt.Println("BUT WHY")

	// otherwise, execute the command
	cmd := exec.Command("bash", "-c", input)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func readInput() (string, error) {
	return reader.ReadString('\n')
}
