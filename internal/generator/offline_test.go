/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package generator

import (
	"testing"
	"unicode/utf8"

	"github.com/lucasloureiror/AegisPass/internal/charsets"
	"github.com/lucasloureiror/AegisPass/internal/cli"
)

func TestGenerateOfflinePassWithCorrectSize(t *testing.T) {
	pwd := cli.Input{
		Size: 7,
	}

	offline := offline{}

	generator := passwordGenerator{
		data: pwd,
		mode: offline,
	}

	charsets.Create(&generator.data)
	password, _, _ := generator.mode.generate(&generator.data)

	got := utf8.RuneCountInString(password)

	if got != pwd.Size {
		t.Errorf("GeneratePass received %d, but returned pass with size %d", 7, got)
	}
}
