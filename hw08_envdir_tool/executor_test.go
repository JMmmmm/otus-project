package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("test positive", func(t *testing.T) {
		var testStdout bytes.Buffer
		stdout = io.Writer(&testStdout)
		stdin = new(bytes.Buffer)

		env := Environment{
			"BAR":   EnvValue{NeedRemove: false, Value: "bar"},
			"EMPTY": EnvValue{NeedRemove: true, Value: ""},
			"FOO":   EnvValue{NeedRemove: false, Value: "   foo\nwith new line"},
		}

		cmd := []string{
			"/bin/bash",
			"testdata/echo.sh",
		}

		returnCode := RunCmd(cmd, env)
		require.Equal(t, 0, returnCode)

		actualResult := testStdout.String()
		expectedResult := "HELLO is ()\nBAR is (bar)\nFOO is (   foo\nwith new line)\nUNSET is ()\nADDED is ()\nEMPTY is" +
			" ()\narguments are \n"
		require.Equal(t, expectedResult, actualResult)
	})
}
