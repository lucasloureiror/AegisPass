/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package generator

import (
	"github.com/lucasloureiror/AegisPass/internal/charsets"
	"github.com/lucasloureiror/AegisPass/internal/config"
	"github.com/lucasloureiror/AegisPass/internal/output"
	"github.com/lucasloureiror/AegisPass/internal/randomclient"
	"github.com/lucasloureiror/AegisPass/internal/shuffle"
)

func Start(pwd *config.Password) {

	generator := GeneratorData{
		behavior: nil,
		password: *pwd,
		random:   nil,
	}

	charsets.Create(pwd)

	if pwd.Flags.Offline {
		generator.behavior = offline{}
		generator.behavior.generate(pwd, nil)
		output.Print(pwd)
		return
	}

	apiResponse := make(chan []string)

	go func() {
		apiResponse <- randomclient.Start(pwd)
	}()

	shuffle.Byte(&pwd.CharSet)

	if pwd.Flags.UseStandard {
		generator.behavior = standard{}
		generator.behavior.generate(pwd, <-apiResponse)
	} else {
		generator.behavior = random{}
		generator.behavior.generate(pwd, <-apiResponse)
	}

	output.Print(pwd)

}
