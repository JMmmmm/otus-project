package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	environment := make(Environment)
	for _, fileInfo := range files {
		fileData, err := os.ReadFile(dir + "/" + fileInfo.Name())
		if err != nil {
			log.Println(err)
			continue
		}

		fileRows := bytes.Split(fileData, []byte("\n"))
		if len(fileRows) == 0 {
			environment[fileInfo.Name()] = EnvValue{NeedRemove: true}
			continue
		}
		fileFormattedRows := bytes.ReplaceAll(fileRows[0], []byte("\u0000"), []byte("\n"))

		envValue := strings.TrimRight(string(fileFormattedRows), "\t")
		envValue = strings.TrimRight(envValue, " ")

		environment[fileInfo.Name()] = EnvValue{NeedRemove: false, Value: envValue}
	}

	return environment, nil
}
