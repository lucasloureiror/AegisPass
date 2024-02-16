/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package generator

import (
	"fmt"

	"github.com/lucasloureiror/AegisPass/internal/cli"
)

type PasswordGeneratorStrategy interface {
	generate(*cli.Input) (string, int, error) //Passing the model for a password ande slice of strings with api generated numbers.
}

type passwordGenerator struct {
	mode      PasswordGeneratorStrategy
	data      cli.Input
	generated []string
	credits   int
}

func (pwd *passwordGenerator) print() {
	if pwd.data.Flags.PrintCredits {
		fmt.Println("API credits remaining:", fmt.Sprint(pwd.credits))
	}
	for _, password := range pwd.generated {
		fmt.Println(password)
	}
}
