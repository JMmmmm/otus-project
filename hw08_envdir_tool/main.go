package main

import (
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

	RunCmd(inputArguments[2:], env)
}
