package security

import "errors"

// ValidateLoginStructure validates login structure
func ValidateLoginStructure(login string) (err error) {
	if len(login) < 4 || len(login) > 12 {
		err = errors.New("Login must have length between 4 and 12 letters")
	}
	return
}
