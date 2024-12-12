package utils

import (
	"fmt"
)

var Verbose bool = false
var Part int = 1

func ParseArgs(args []string) {
	for i := 0; i < len(args); i++ {
		if args[i] == "--part-two" || args[i] == "-t" {
			Part = 2
		} else if args[i] == "--verbose" || args[i] == "-v" {
			Verbose = true
		}
	}
}
func Log(a ...any) (n int, err error) {
	if !Verbose {
		return 0, nil
	}
	return fmt.Println(a...)
}
func LogF(format string, a ...any) (n int, err error) {
	if !Verbose {
		return 0, nil
	}
	return fmt.Printf(format, a...)
}
