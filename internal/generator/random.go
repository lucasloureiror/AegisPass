/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package generator

import (
	"strconv"

	"github.com/lucasloureiror/AegisPass/internal/config"
	"github.com/lucasloureiror/AegisPass/internal/shuffle"
)

type random struct{}

func (random) generate(pwd *config.Password, randomIndex []string) {

	var index int
	pwdLen := pwd.Size - len(pwd.Generated)

	for i := 0; i < pwdLen; i++ {
		index, _ = strconv.Atoi(randomIndex[i])
		pwd.Generated = pwd.Generated + string(pwd.CharSet[index])
	}
	shuffle.String(&pwd.Generated)

}
