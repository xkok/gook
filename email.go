package gook

import (
	"errors"
	"regexp"
)

// Email is a type for email addresses.
// # example
//
//	gook.Email("HtH8O@example.com").Ok() // true
//	gook.Email("example.com").Error() // not valid email
type Email string

var rxMail = regexp.MustCompile(`^(?:[[:alnum:]]+[[:alnum:]\-\.]+[[:alnum:]])+@(?:[[:alnum:]]+[[:alnum:]\-\.]+[[:alnum:]])+\.(?:[[:alpha:]]{2,6})$`)

// Error checks if the given email is valid.
//
// It takes no parameters.
// It returns an error.
func (e Email) Error() (err error) {
	if !rxMail.MatchString(string(e)) {
		return errors.New("not valid email")
	}
	return
}

// Ok checks if the Email is valid.
//
// It returns a boolean indicating whether the Email is valid or not.
func (e Email) Ok() bool {
	return rxMail.MatchString(string(e))
}
