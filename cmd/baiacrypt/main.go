package baiacrypt

import (
	"fmt"
	"os"

	"github.com/brunocampos-ssa/baiacrypt/crypto"

	"github.com/brunocampos-ssa/baiacrypt/prompt"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage:")
		fmt.Println("  encrypt <input.txt> <output.enc>")
		fmt.Println("  decrypt <input.enc> <output.txt>")
		return
	}

	mode := os.Args[1]
	inputFile := os.Args[2]
	outputFile := os.Args[3]

	var password []byte
	var err error

	switch mode {
	case "encrypt":
		password, err = prompt.PromptPassword(true) // with confirmation
	case "decrypt":
		password, err = prompt.PromptPassword(false)
	default:
		fmt.Println("Invalid mode:", mode)
		return
	}
	if err != nil {
		fmt.Println("Password error:", err)
		return
	}

	inputData, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	var outputData []byte
	switch mode {
	case "encrypt":
		outputData, err = crypto.Encrypt(inputData, password)
	case "decrypt":
		outputData, err = crypto.Decrypt(inputData, password)
	}
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(outputFile, outputData, 0644); err != nil {
		panic(err)
	}

	fmt.Printf("%s successful: %s\n", mode, outputFile)
}
