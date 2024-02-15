package generator

import (
	"github.com/lucasloureiror/AegisPass/internal/cli"
)

type help struct{}

func (h help) generate(pwd *cli.Input, apiResponse []string) string {
	return ""
}

func (h help) print() string {

	helpMessage := `
AegisPass v1.2.2
AegisPass source code is licensed under the Mozilla Public License 2.0 and is available at github.com/lucasloureiror/AegisPass
Usage: aegis [password_length] [options]

Arguments:
  password_length  The length of the password to be generated (default: 10)

Options:
  --offline 	   Generate random or numeric passwords without using random.org API
  --numeric        Password with numbers only (default: password with length 10 if not specified)
  --standard       Generate password with one upper case, one number and one special character at least.
  --credits        Print random.org API credits to the user after generating a password
  --help           Help the user to use the CLI tool

Example:
  aegis 12 --numeric
  aegis 8
  aegis 10 --standard
  aegis 10 --credits
  aegis help
`
	return helpMessage
}
