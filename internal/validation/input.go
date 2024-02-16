/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package validation

import (
	"errors"
	"github.com/lucasloureiror/AegisPass/internal/cli"
	"github.com/lucasloureiror/AegisPass/internal/output"
	"os"
	"strconv"
)

func Start(input *cli.Input) error {
	var fetchErr error

	if input.Flags.NeedHelp {
		return nil
	}

	input.Size, fetchErr = fetchSize(os.Args)
	if fetchErr != nil {
		return fetchErr
	}

	sizeError := sizeCheck(input.Size)
	if sizeError != nil {
		return sizeError
	}

	return nil
}

func sizeCheck(size int) error {

	if size <= 3 || size > 35 {
		return errors.New("password size must be bigger than 3 and smaller than 35")

	}
	return nil
}

func fetchSize(args []string) (int, error) {

	var convertErr error
	var size int

	if len(args) < 2 || (args)[1] == "" {
		size = 15
	} else {
		size, convertErr = strconv.Atoi((args)[1])
	}

	if convertErr != nil {
		warning := "Password length not detected, generating password with default length(15)"
		size = 15
		output.PrintWarning(warning)
		return size, nil
	}

	return size, nil

}
