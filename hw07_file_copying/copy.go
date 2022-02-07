package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type WriteCounter struct {
	TotalByte      uint64
	CurrentPercent uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.CurrentPercent += (uint64(n) / wc.TotalByte) * 100
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rCopying... %d%s complete", wc.CurrentPercent, "%")
}

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	var err error
	inputFile, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	stat, err := inputFile.Stat()
	if err != nil {
		return err
	}
	inputFileSize := stat.Size()
	if offset > inputFileSize {
		return ErrOffsetExceedsFileSize
	}

	outputFile, err := os.Create(toPath)
	if err != nil {
		return ErrUnsupportedFile
	}

	_, err = inputFile.Seek(offset, 0)
	if err != nil {
		return err
	}

	counter := &WriteCounter{TotalByte: uint64(getExpectedOutputFileSize(inputFileSize, offset, limit))}
	if limit == 0 || (limit+offset) > inputFileSize {
		_, err = io.Copy(outputFile, io.TeeReader(inputFile, counter))
	} else {
		_, err = io.CopyN(outputFile, io.TeeReader(inputFile, counter), limit)
	}
	if err != nil {
		return err
	}

	return nil
}

func getExpectedOutputFileSize(inputFileSize, offset, limit int64) int64 {
	if limit == 0 || (limit+offset) > inputFileSize {
		return inputFileSize - offset
	}

	return limit - offset
}
