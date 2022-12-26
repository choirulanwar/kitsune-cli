package cmd

import (
	"fmt"
	"os"
	"strings"
)

func GetWorkingDirectory() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	return cwd
}

func Replace(input, from, to string) string {
	return strings.Replace(input, from, to, -1)
}

func RemoveDash(s string) string {
	return strings.Replace(s, "-", "", -1)
}

func RemoveSpace(s string) string {
	return strings.Replace(s, " ", "", -1)
}

func Add(a, b int) int {
	return a + b
}
