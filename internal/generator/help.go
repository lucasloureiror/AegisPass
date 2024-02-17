/*
Copyright 2024 github.com/lucasloureiror/AegisPass maintainers

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
package generator

import (
	"github.com/lucasloureiror/AegisPass/internal/cli"
)

type help struct{}

func (h help) generate(pwd *cli.Input) (string, int, error) {
	helpMessage := `
AegisPass v2.0.0
AegisPass source code is licensed under the Mozilla Public License 2.0 and is available at github.com/lucasloureiror/AegisPass
Usage: aegis [options] [password_length]

Arguments:
  password_length  The length of the password to be generated (default: 15)

Options:
  --bulk 		   Specify the number of passwords to be generated (default: 1)
  --numeric        Generate a numeric password
  --standard       Generate password with one upper case, one number and one special character at least.
  --online 	   	   Generate random or numeric passwords with random.org web service
  --credits        Print random.org web service credits to the user after generating a password
  --help           Print the help message to the user

Example:
  aegis --numeric 12
  aegis 8
  aegis --standard 10
  aegis --online 15
  aegis --online --credits 10
  aegis --help
`
	return helpMessage, -1, nil
}
