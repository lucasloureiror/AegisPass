/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package shuffle

import ()

func FisherYates(arr []byte, length int) []byte {
	shuffled := make([]byte, length)
	for i := 0; i < length; i++ {
		j := randomInt(len(arr))
		shuffled[i] = arr[j]
	}

	return shuffled
}
