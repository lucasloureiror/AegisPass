/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package generator

import (
	"github.com/lucasloureiror/AegisPass/internal/cli"
	"github.com/lucasloureiror/AegisPass/internal/shuffle"
)

type random struct{}

func (random) generate(input *cli.Input) (string, int, error) {
	shuffle.Byte(&input.CharSet)
	password := shuffle.BuildString(input.CharSet, input.Size)
	shuffle.String(&password)
	return password, -1, nil
}
