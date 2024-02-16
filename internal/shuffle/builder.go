/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package shuffle

// This function builds a string with CSPRNG using the original charset and the length of the password
func BuildString(arr []byte, length int) string {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		j := randomInt(len(arr))
		result[i] = arr[j]
	}

	return string(result)
}
