package commands

import (
	"errors"
	"os"
	"time"
)

func Touch(argv []string) error {
	if len(argv) < 2 {
		return errors.New("missing file operand")
	}
	for _, filename := range argv[1:] {
		if _, err := os.Stat(filename); err != nil {
			file, err := os.Create(filename)
			if err != nil {
				return err
			}
			if err := file.Close(); err != nil {
				return err
			}
		} else {
			currentTime := time.Now().Local()
			if err := os.Chtimes(filename, currentTime, currentTime); err != nil {
				return err
			}
		}
	}
	return nil
}
