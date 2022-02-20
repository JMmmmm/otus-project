package main

import (
	"io"
	"os"
	"os/exec"
)

const (
	cmdErrorCode = 1
	cmdOkCode    = 0
)

var (
	stdout io.Writer
	stdin  io.Reader
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for fileName, fileEnvValue := range env {
		_, ok := os.LookupEnv(fileName)
		if ok {
			err := os.Unsetenv(fileName)
			if err != nil {
				return cmdErrorCode
			}
		}
		if !fileEnvValue.NeedRemove {
			err := os.Setenv(fileName, fileEnvValue.Value)
			if err != nil {
				return cmdErrorCode
			}
		}
	}

	cmdCommand := cmd[0]
	cmdArguments := cmd[1:]
	command := exec.Command(cmdCommand, cmdArguments...)
	command.Stdout = stdout
	command.Stdin = stdin

	if err := command.Run(); err != nil {
		return cmdErrorCode
	}

	return cmdOkCode
}
