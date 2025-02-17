package utils

import "errors"

func PasswordValidation(password string) error {

	if IsEmpty(password) {
		return errors.New("Password is empty!")
	}
	if !HasEnoughLetters(password, 8) {
		return errors.New("Password must contain 8 or more characters!")
	}

	return nil
}
