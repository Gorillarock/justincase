//Package with helpful regex tools for Go
package justincase

import (
	"regexp"
	"unicode"
)

// Returns regex string that will accept any string which matches the given param string (regardless of case)
func GenerateCaseInsensitiveRegex(param string) *regexp.Regexp {
	if param == "" {
		return regexp.MustCompile("^$")
	}
	var regex []rune
	regex = append(regex, '(')
	for _, char := range param {
		iterRunes := []rune{'(', '?', ':', unicode.ToUpper(char), '|', unicode.ToLower(char), ')', '+'}
		regex = append(regex, iterRunes...)
	}
	regex = append(regex, ')')
	return regexp.MustCompile(string(regex))
}
