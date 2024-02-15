/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package shuffle

import (
	"bytes"
	"testing"

	"github.com/lucasloureiror/AegisPass/internal/cli"
)

func TestShuffle(t *testing.T) {

	pwd := cli.Input{
		CharSet: []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*"),
	}

	want := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*")

	Byte(&pwd.CharSet)

	got := pwd.CharSet
	if bytes.Equal(got, want) {
		t.Error("Array wasn't shuffled")
	}

}
