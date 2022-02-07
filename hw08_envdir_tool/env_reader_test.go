package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("test positive", func(t *testing.T) {
		result, err := ReadDir("testdata/env")

		env := Environment{
			"BAR":   EnvValue{NeedRemove: false, Value: "bar"},
			"EMPTY": EnvValue{NeedRemove: false, Value: ""},
			"FOO":   EnvValue{NeedRemove: false, Value: "   foo\nwith new line"},
			"HELLO": EnvValue{NeedRemove: false, Value: "\"hello\""},
			"UNSET": EnvValue{NeedRemove: false, Value: ""},
		}

		require.NoError(t, err)
		require.Equal(t, env, result)
	})
}
