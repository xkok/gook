package gook

import (
	"errors"
	"fmt"
	"regexp"
)

// Password is a type for password.
// # example
//
//	gook.Password("password").Error() // error
//	gook.Password("5uperP@ssw0rd").Ok() // true
type Password string

var passMin int = 8
var upper = regexp.MustCompile(`[[:upper:]]+`)
var lower = regexp.MustCompile(`[[:lower:]]+`)
var punct = regexp.MustCompile(`[[:punct:]]+`)
var digit = regexp.MustCompile(`[[:digit:]]+`)

// Set minimum character
func (p Password) Min(integer int) Password {
	passMin = integer
	return p
}

// Error checks if a given password is valid.
//
// It takes no parameters.
// It returns an error if the password does not meet the required criteria.
func (p Password) Error() (err error) {
	if len(string(p)) < int(passMin) {
		err = fmt.Errorf("password should contains %d or more character", passMin)
	}
	if !upper.MatchString(string(p)) {
		err = errors.New("password should contain one or more uppercase character")
	}
	if !lower.MatchString(string(p)) {
		err = errors.New("password should contain one or more lowercase character")
	}
	if !punct.MatchString(string(p)) {
		err = errors.New("password should contain one or more special character")
	}
	if !digit.MatchString(string(p)) {
		err = errors.New("password should contain one or more digit character")
	}
	return
}

// Ok checks if the password meets the requirements.
//
// It checks if the password contains at least one uppercase letter,
// one lowercase letter, one punctuation character, and one digit.
// Returns true if the password meets the requirements, false otherwise.
func (p Password) Ok() bool {
	if len(string(p)) < passMin {
		return false
	}
	if !upper.MatchString(string(p)) {
		return false
	}
	if !lower.MatchString(string(p)) {
		return false
	}
	if !punct.MatchString(string(p)) {
		return false
	}
	if !digit.MatchString(string(p)) {
		return false
	}
	return true
}
