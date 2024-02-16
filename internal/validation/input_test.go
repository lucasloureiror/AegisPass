/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package validation

import (
	"testing"
)

func TestGenerateInvalidPassSizes(t *testing.T) {
	size := -1
	got := sizeCheck(size)

	if got == nil {
		t.Errorf("Expected error, but received none with password size %d", size)
	}

	size = 3

	got = sizeCheck(size)

	if got == nil {
		t.Errorf("Expected error, but received none with password size %d", size)
	}

	size = 36

	got = sizeCheck(size)

	if got == nil {
		t.Errorf("Expected error, but received none with password size %d", size)
	}
}

func TestFetchSizeWithoutValue(t *testing.T) {

	args := "teste"
	defaultSize := 15

	size := fetchSize(args)

	if size != defaultSize {
		t.Errorf("Expected size %d, but received size %d", defaultSize, size)
	}

}
