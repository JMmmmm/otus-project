package hw02unpackstring

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidString = errors.New("invalid string")
	digitCheck       = regexp.MustCompile(`^[0-9]$`)
)

func Unpack(incomingString string) (string, error) {
	if incomingString == "" {
		return "", nil
	}

	var result strings.Builder
	var previousIncomingStringPart string
	var isPreviousPartScreened bool

	for key, incomingPart := range incomingString {
		incomingStringPart := string(incomingPart)
		isScreening := previousIncomingStringPart == `\` && !isPreviousPartScreened
		switch {
		case isScreening:
			previousIncomingStringPart = incomingStringPart
		case digitCheck.MatchString(incomingStringPart):
			if key == 0 || (previousIncomingStringPart == "" && !isPreviousPartScreened) {
				return "", ErrInvalidString
			}
			incomingIntPart := int(incomingPart - '0')
			if incomingIntPart == 0 {
				previousIncomingStringPart = ""
				continue
			}

			var addition strings.Builder
			for i := 0; i < incomingIntPart; i++ {
				addition.WriteString(previousIncomingStringPart)
			}
			result.WriteString(addition.String())
			previousIncomingStringPart = ""
		default:
			result.WriteString(previousIncomingStringPart)
			previousIncomingStringPart = incomingStringPart
		}
		isPreviousPartScreened = isScreening
	}
	result.WriteString(previousIncomingStringPart)

	return result.String(), nil
}
