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

type offline struct{}

func (offline) generate(pwd *cli.Input) (string, int, error) {
	shuffle.Byte(&pwd.CharSet)
	pwdByte := shuffle.FisherYates(pwd.CharSet, pwd.Size)
	return string(pwdByte), -1, nil
}
