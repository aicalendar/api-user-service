package utils

import "errors"

func EmailValidation(email string) error {

	if IsEmpty(email) {
		return errors.New("User name is empty!")
	}
	if !IsValidEmail(email) {
		return errors.New("User name must contain 3 or more characters!")
	}

	return nil
}
