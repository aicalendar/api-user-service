package utils

import "errors"

func UsernameValidation(password string) error {

	if IsEmpty(password) {
		return errors.New("User name is empty!")
	}
	if !HasEnoughLetters(password, 3) {
		return errors.New("User name must contain 3 or more characters!")
	}
	if !isAlphanumeric(password) {
		return errors.New("User name canno't contain special characters!")
	}

	return nil
}
