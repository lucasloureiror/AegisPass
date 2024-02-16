/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package main

import (
	"github.com/lucasloureiror/AegisPass/internal/cli"
	"github.com/lucasloureiror/AegisPass/internal/generator"
	"github.com/lucasloureiror/AegisPass/internal/output"
	"github.com/lucasloureiror/AegisPass/internal/validation"
)

func main() {
	var input cli.Input

	cli.ParseFlags(&input)

	err := validation.Start(&input)

	if err != nil {
		output.PrintError(err.Error())
		return
	}

	mode := generator.ReturnGeneratorMode(&input)

	generator.Start(input, mode)

}
