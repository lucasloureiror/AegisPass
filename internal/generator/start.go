/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package generator

import (
	"github.com/lucasloureiror/AegisPass/internal/charsets"
	"github.com/lucasloureiror/AegisPass/internal/cli"
)

func Start(input cli.Input, mode PasswordGeneratorStrategy) {

	generator := passwordGenerator{
		mode: mode,
		data: input,
	}

	charsets.Create(&generator.data)

	generator.generated, generator.credits, _ = generator.mode.generate(&generator.data)

	generator.print()

}
