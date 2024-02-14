/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package generator

import "github.com/lucasloureiror/AegisPass/internal/config"

type generatorInterface interface {
	generate(*config.Password, []string) //Passing the model for a password ande slice of strings with api generated numbers.
}

type GeneratorData struct {
	behavior generatorInterface
	password config.Password
	random   []string
}
