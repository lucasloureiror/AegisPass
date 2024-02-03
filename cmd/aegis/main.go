/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package main

import (
	"github.com/lucasloureiror/AegisPass/internal/config"
	"github.com/lucasloureiror/AegisPass/internal/generator"
	"github.com/lucasloureiror/AegisPass/internal/validation"
)

func main() {
	var pwd config.Password

	config.ParseFlags(&pwd.Flags)

	validation.Start(&pwd)

	generator.Start(&pwd)

}
