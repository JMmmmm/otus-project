package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	errorValidationLength = errors.New("incorrect length")
	errorValidationMin    = errors.New("min value is exceeded")
	errorValidationMax    = errors.New("max value is exceeded")
	errorValidationIn     = errors.New("value is not exists in list")
	errorValidationRegex  = errors.New("incorrect value by regex")
)

var (
	errorInvalidValueType      = errors.New("invalid value type")
	errorInvalidFormatField    = errors.New("invalid format field")
	errorInvalidValidationRule = errors.New("invalid validation rule")
	errorInvalidLenRuleType    = errors.New("invalid len rule type")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	errorsCount := len(v)
	formattedErrors := make([]string, errorsCount)
	for key, validationError := range v {
		formattedErrors[key] = fmt.Sprintf("field: %v error: %v", validationError.Field, validationError.Err)
	}

	return "Validation errors: " + strings.Join(formattedErrors, "\n")
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("field: %v; error: %v", v.Field, v.Err)
}

func Validate(v interface{}) error {
	reflectValue := reflect.ValueOf(v)
	if reflectValue.Kind() != reflect.Struct {
		return errorInvalidValueType
	}
	reflectType := reflectValue.Type()

	var validationErrors ValidationErrors
	for i := 0; i < reflectValue.NumField(); i++ {
		fieldStruct := reflectType.Field(i)
		structTagValidate := fieldStruct.Tag.Get("validate")
		if structTagValidate == "" {
			continue
		}
		validations := strings.Split(structTagValidate, "|")

		for _, validation := range validations {
			validationRule := strings.Split(validation, ":")
			if len(validationRule) != 2 {
				return errorInvalidFormatField
			}
			commonErr := validateByRule(validationRule[0], validationRule[1], reflectValue.Field(i))
			if commonErr == nil {
				continue
			}
			validationError, ok := commonErr.(ValidationError)
			if !ok {
				return fmt.Errorf("validate program error of field - \"%s\" : %w", fieldStruct.Name, commonErr)
			}
			validationError.Field = fieldStruct.Name
			validationErrors = append(validationErrors, validationError)
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

func validateByRule(validationRuleName string, validationRuleValue string, fieldValue reflect.Value) error {
	if fieldValue.IsZero() {
		return nil
	}
	switch validationRuleName {
	case "len":
		return validateLen(validationRuleValue, fieldValue)
	case "min":
		return validateMin(validationRuleValue, fieldValue)
	case "max":
		return validateMax(validationRuleValue, fieldValue)
	case "in":
		return validateIn(validationRuleValue, fieldValue)
	case "regex":
		return validateRegex(validationRuleValue, fieldValue)
	}

	return errorInvalidValidationRule
}

func validateLen(validationRuleValue string, fieldValue reflect.Value) error {
	ruleLen, err := strconv.Atoi(validationRuleValue)
	if err != nil {
		return err
	}

	var valueLen int
	switch fieldValue.Kind() {
	case reflect.String:
		valueLen = len(fieldValue.String())
	case reflect.Slice:
		valueLen = len(fieldValue.Interface().([]string))
	default:
		return errorInvalidLenRuleType
	}

	if ruleLen != valueLen {
		return ValidationError{Err: errorValidationLength}
	}

	return nil
}

func validateMin(validationRuleValue string, fieldValue reflect.Value) error {
	ruleMin, err := strconv.Atoi(validationRuleValue)
	if err != nil {
		return err
	}
	if int64(ruleMin) > fieldValue.Int() {
		return ValidationError{Err: errorValidationMin}
	}

	return nil
}

func validateMax(validationRuleValue string, fieldValue reflect.Value) error {
	ruleMax, err := strconv.Atoi(validationRuleValue)
	if err != nil {
		return err
	}
	if int64(ruleMax) < fieldValue.Int() {
		return ValidationError{Err: errorValidationMax}
	}

	return nil
}

func validateIn(validationRuleValue string, fieldValue reflect.Value) error {
	ruleIn := strings.Split(validationRuleValue, ",")

	var value string
	if fieldValue.Kind() == reflect.Int {
		value = strconv.Itoa(int(fieldValue.Int()))
	} else {
		value = fieldValue.String()
	}
	for _, availableElem := range ruleIn {
		if availableElem == value {
			return nil
		}
	}

	return ValidationError{Err: errorValidationIn}
}

func validateRegex(validationRuleValue string, fieldValue reflect.Value) error {
	ruleRegex := regexp.MustCompile(validationRuleValue)
	isMatchRegex := ruleRegex.MatchString(fieldValue.String())

	if isMatchRegex {
		return nil
	}

	return ValidationError{Err: errorValidationRegex}
}
