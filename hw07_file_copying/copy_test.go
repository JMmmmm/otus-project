package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	tests := []struct {
		expectedOutput string
		offset, limit  int64
	}{
		{expectedOutput: "testdata/out_offset0_limit0", offset: 0, limit: 0},
		{expectedOutput: "testdata/out_offset0_limit10", offset: 0, limit: 10},
		{expectedOutput: "testdata/out_offset0_limit1000", offset: 0, limit: 1000},
		{expectedOutput: "testdata/out_offset0_limit10000", offset: 0, limit: 10000},
		{expectedOutput: "testdata/out_offset100_limit1000", offset: 100, limit: 1000},
		{expectedOutput: "testdata/out_offset6000_limit1000", offset: 6000, limit: 1000},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.expectedOutput, func(t *testing.T) {
			resultOutput := tc.expectedOutput + "test.txt"
			defer func() {
				err := os.Remove(resultOutput)
				if err != nil {
					t.Fatalf("failed removing: %s", err)
				}
			}()

			err := Copy("testdata/input.txt", resultOutput, tc.offset, tc.limit)
			if err != nil {
				t.Fatalf("failed test: %s", err)
			}
			resultBytes, err := ioutil.ReadFile(resultOutput)
			if err != nil {
				t.Fatalf("failed reading: %s", err)
			}
			expectedBytes, err := ioutil.ReadFile(tc.expectedOutput + ".txt")
			if err != nil {
				t.Fatalf("failed reading: %s", err)
			}

			require.Equal(t, expectedBytes, resultBytes)
		})
	}

	t.Run("test offset exceeds file size", func(t *testing.T) {
		err := Copy("testdata/input.txt", "testdata/input-test.txt", 10000, 0)

		require.Equal(t, ErrOffsetExceedsFileSize, err)
	})
}
