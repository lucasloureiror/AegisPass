/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package charsets

import "github.com/lucasloureiror/AegisPass/internal/config"

func Create(pwd *config.Password) {

	var (
		nums = []byte("0123456789")
		full = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*")
	)

	if pwd.Flags.UseOnlyNums {
		pwd.CharSet = append(pwd.CharSet, nums...)
		return
	} else {
		pwd.CharSet = full

	}

}
