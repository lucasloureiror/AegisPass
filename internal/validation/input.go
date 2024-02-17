/*
Copyright 2024 lucasloureiror

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
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
		warning := "Password length not detected, generating password with default length(15), use flag --help for more details"
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
