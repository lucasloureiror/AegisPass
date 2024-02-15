/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package generator

import (
	"github.com/lucasloureiror/AegisPass/internal/cli"
	"github.com/lucasloureiror/AegisPass/internal/randomclient"
	"github.com/lucasloureiror/AegisPass/internal/shuffle"
	"strconv"
	"sync"
)

type standard struct{}

func (standard) generate(input *cli.Input) (string, int, error) {

	var wg sync.WaitGroup
	var credits int
	var randomIndex []string
	wg.Add(5)
	upper := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	special := []byte("!@#$%&*")
	nums := []byte("0123456789")

	go func() {
		defer wg.Done()
		shuffle.Byte(&upper)
	}()
	go func() {
		defer wg.Done()
		shuffle.Byte(&special)
	}()

	go func() {
		defer wg.Done()
		shuffle.Byte(&nums)
	}()

	go func() {
		defer wg.Done()
		randomIndex, credits, _ = randomclient.Start(input)
	}()

	go func() {
		defer wg.Done()
		shuffle.Byte(&input.CharSet)
	}()

	wg.Wait()

	generated := string(upper[0]) + string(special[0]) + string(nums[0])

	var index int
	pwdLen := input.Size - len(generated)

	for i := 0; i < pwdLen; i++ {
		index, _ = strconv.Atoi(randomIndex[i])
		generated = generated + string(input.CharSet[index])
	}
	shuffle.String(&generated)

	return generated, credits, nil
}
