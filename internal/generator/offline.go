/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package generator

import (
	"github.com/lucasloureiror/AegisPass/internal/config"
	"github.com/lucasloureiror/AegisPass/internal/shuffle"
)

type offline struct{}

func (offline) generate(pwd *config.Password, _ []string) {
	shuffle.Byte(&pwd.CharSet)
	pwdByte := shuffle.FisherYates(pwd.CharSet, pwd.Size)
	pwd.Generated = string(pwdByte)
}
