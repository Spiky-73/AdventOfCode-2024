package utils

import (
	"fmt"
	"strconv"
)

var Verbose bool = false
var Part int = 1

func ParseArgs(args []string) {
	for i := 0; i < len(args); i++ {
		if args[i] == "-p" {
			Part, _ = strconv.Atoi((args[i+1]))
			i++
		} else if args[i] == "-v" {
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
