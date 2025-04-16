package prompt

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/term"
)

// PromptPassword reads a password from terminal and optionally confirms it.
func PromptPassword(confirm bool) ([]byte, error) {
	return readPasswordWithConfirm(confirm, os.Stdin)
}

// readPasswordWithConfirm allows injecting the input source (used in tests)
func readPasswordWithConfirm(confirm bool, in *os.File) ([]byte, error) {
	fmt.Print("Enter password: ")
	password, err := term.ReadPassword(int(in.Fd()))
	fmt.Println()
	if err != nil {
		return nil, err
	}

	if confirm {
		fmt.Print("Confirm password: ")
		confirmPassword, err := term.ReadPassword(int(in.Fd()))
		fmt.Println()
		if err != nil {
			return nil, err
		}
		if string(password) != string(confirmPassword) {
			return nil, errors.New("passwords do not match")
		}
	}

	return password, nil
}
