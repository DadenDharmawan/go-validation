package govalidation

import (
	"regexp"
)

func Lambda(name string) bool {
	if name != "" {
		regex := regexp.MustCompile("[0-9]+")
		ismatch := regex.MatchString(name)
		if !ismatch {
			return true
		}
	}
	return false
}
