/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package shuffle

import (
	"crypto/rand"
	"math/big"
)

func randomInt(max int) int {
	maxBigInt := big.NewInt(int64(max))
	val, err := rand.Int(rand.Reader, maxBigInt)
	if err != nil {
		panic(err)
	}
	return int(val.Int64())
}
