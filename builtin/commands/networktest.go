package commands

import "net/http"

func Networktest(argv []string) error {
	_, err := http.Get("http://www.google.com")
	return err
}
