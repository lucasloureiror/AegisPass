/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package shuffle

import (
	"testing"
)

func TestShuffleString(t *testing.T) {

	original := "TestString1235)!@#"

	shuffled := original

	String(&shuffled)

	if original == shuffled {
		t.Error("Original string was: ", original, "Shuffled string was: ", shuffled, "Strings are equal, they shouldn't be.")
	}

}
