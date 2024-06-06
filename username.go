package gook

import "fmt"

type UserName string

var unameMin int = 5

// Set minimum character
func (u UserName) Min(integer int) UserName {
	unameMin = integer
	return u
}

func (u UserName) Error() (err error) {
	if len(string(u)) < unameMin {
		err = fmt.Errorf("username should contains %d or more character", unameMin)
	}
	return
}

func (u UserName) Ok() bool {
	return len(string(u)) > unameMin
}
