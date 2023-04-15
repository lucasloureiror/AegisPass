package output

import (
	"fmt"
	"os"

	"github.com/lucasloureiror/AegisPass/internal/config"
)

func Print(pwd *config.Password) {
	if pwd.Flags.PrintCredits {
		fmt.Println(pwd.Generated)
		fmt.Printf("API credits remaining: %d", pwd.APICredit)
	} else {
		fmt.Println(pwd.Generated)
	}
}

func PrintError(err string) {
	fmt.Println("Error: ", err)
}

func PrintHelp() {

	helpMessage := `
Usage: aegis [password_length] [options]

Arguments:
  password_length  The length of the password to be generated (default: 10)

Options:
  --numeric        Password with numbers only (default: password with length 10 if not specified)
  --credits        Print random.org API credits to the user after generating a password
  --help           Help the user to use the CLI tool

Example:
  aegis 12 --numeric
  aegis 8
  aegis 10 --credits
  aegis help
`
	fmt.Print(helpMessage)
	os.Exit(0)
}
