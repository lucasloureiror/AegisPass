/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package validation

import (
	"errors"
	"flag"
	"strconv"

	"github.com/lucasloureiror/AegisPass/internal/cli"
	"github.com/lucasloureiror/AegisPass/internal/output"
)

func Start(input *cli.Input) error {
	var err error

	if input.Flags.NeedHelp {
		return nil
	}

	input.Size = fetchSize(flag.Arg(0))

	err = sizeCheck(input.Size)
	if err != nil {
		return err
	}

	err = numberOfPasswordsCheck(input.NumberOfPasswords)

	if err != nil {
		return err
	}

	return nil
}

func sizeCheck(size int) error {

	if size <= 3 || size > 35 {
		return errors.New("password size must be bigger than 3 and smaller than 35")

	}
	return nil
}

func fetchSize(args string) int {
	size, convertErr := strconv.Atoi(args)

	if convertErr != nil {
		warning := "Password length not detected, generating password with default length(15)"
		size = 15
		output.PrintWarning(warning)
	}
	return size

}

func numberOfPasswordsCheck(number int) error {
	if number < 1 || number > 10 {
		return errors.New("number of passwords must be bigger than 0 and smaller than 10")
	}
	return nil
}
