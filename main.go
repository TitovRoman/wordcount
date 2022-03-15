package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    if len(os.Args) == 1 {
		fmt.Fprint(os.Stderr, "Pass the second parameter\n")
		os.Exit(1)
	}
	inputString := os.Args[1]
	cnt := len(strings.Fields(inputString))
	fmt.Println(cnt)
}
