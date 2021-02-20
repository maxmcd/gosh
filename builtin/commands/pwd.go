package commands

import (
	"fmt"
	"os"
)

func Pwd(argv []string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(dir)
	return err
}
