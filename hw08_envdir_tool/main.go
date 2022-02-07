package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	inputArguments := os.Args

	stdout = os.Stdout
	stdin = os.Stdin
	env, err := ReadDir(inputArguments[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(RunCmd(inputArguments[2:], env))
}
