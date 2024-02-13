/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
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

func PrintWarning(warn string) {
	fmt.Println("Warning: ", warn)
}

func PrintHelp() {

	helpMessage := `
AegisPass v1.2.1
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
	fmt.Print(helpMessage)
	os.Exit(0)
}
