package utils

import "errors"

func UsernameValidation(name string) error {

	if IsEmpty(name) {
		return errors.New("User name is empty!")
	}
	if !HasEnoughLetters(name, 3) {
		return errors.New("User name must contain 3 or more characters!")
	}
	if !isAlphanumeric(name) {
		return errors.New("User name canno't contain special characters!")
	}

	return nil
}
