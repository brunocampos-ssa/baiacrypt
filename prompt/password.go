package prompt

import (
	"crypto/subtle"
	"errors"
	"fmt"
	"os"

	"golang.org/x/term"
)

func PromptPassword(confirm bool) ([]byte, error) {
	fmt.Print("Enter password: ")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		return nil, err
	}

	if confirm {
		fmt.Print("Confirm password: ")
		confirmPassword, err := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println()
		if err != nil {
			return nil, err
		}
		if subtle.ConstantTimeCompare(password, confirmPassword) != 1 {
			return nil, errors.New("passwords do not match")
		}
	}
	return password, nil
}
