/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package shuffle

func String(pwd *string) {

	pwdBytes := []byte(*pwd)
	strSize := len(pwdBytes) - 1

	for i := strSize; i > 0; i-- {
		j := randomInt(strSize)
		pwdBytes[i], pwdBytes[j] = pwdBytes[j], pwdBytes[i]
	}

	*pwd = string(pwdBytes)

}
