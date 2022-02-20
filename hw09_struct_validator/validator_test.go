package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: User{
				ID:     "rrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr",
				Age:    2,
				Role:   "stuff",
				Phones: []string{"44444", "3223"},
			},
			expectedErr: ValidationErrors{
				ValidationError{Err: errorValidationLength},
				ValidationError{Err: errorValidationMin},
				ValidationError{Err: errorValidationLength},
			},
		},
		{
			in: User{
				Role: "stufffff",
				meta: []byte{33, 33},
			},
			expectedErr: ValidationErrors{
				ValidationError{Err: errorValidationIn},
			},
		},
		{
			in: App{
				Version: "fghjk",
			},
			expectedErr: nil,
		},
		{
			in: App{
				Version: "fghjk1",
			},
			expectedErr: ValidationErrors{
				ValidationError{Err: errorValidationLength},
			},
		},
		{
			in: Token{
				Header: []byte{22, 33},
			},
			expectedErr: nil,
		},
		{
			in:          []int{1, 5},
			expectedErr: errorInvalidValueType,
		},
		{
			in: Response{
				Code: 200,
				Body: "rrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr",
			},
			expectedErr: nil,
		},
		{
			in: Response{
				Code: 100,
				Body: "rrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr",
			},
			expectedErr: ValidationErrors{
				ValidationError{Err: errorValidationIn},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			result := Validate(tt.in)
			if result == nil && tt.expectedErr == nil {
				return
			}
			validationErrors, validationErrorsOk := result.(ValidationErrors)
			expectedValidationErrors, expectedValidationErrorsOk := tt.expectedErr.(ValidationErrors)
			if !validationErrorsOk || !expectedValidationErrorsOk {
				if !errors.Is(result, tt.expectedErr) {
					t.Errorf("not equal program errors:\n expected - %q;\n actual - %q", tt.expectedErr, result)
				}
				return
			}
			for key, validationError := range validationErrors {
				expectedValidationError := expectedValidationErrors[key]
				require.Error(t, expectedValidationError)
				if !errors.Is(validationError.Err, expectedValidationError.Err) {
					t.Errorf("not equal validation errors:\n expected - %q;\n actual - %q",
						expectedValidationError.Err,
						result)
				}
			}

			_ = tt
		})
	}
}
