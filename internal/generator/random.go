/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package generator

import (
	"strconv"
	"sync"

	"github.com/lucasloureiror/AegisPass/internal/cli"
	"github.com/lucasloureiror/AegisPass/internal/randomclient"
	"github.com/lucasloureiror/AegisPass/internal/shuffle"
)

type random struct{}

func (random) generate(input *cli.Input) (string, int, error) {

	var randomIndex []string
	var wg sync.WaitGroup
	var credits int

	wg.Add(2)
	shuffle.Byte(&input.CharSet)

	go func() {
		defer wg.Done()
		randomIndex, credits, _ = randomclient.Start(input)
	}()

	go func() {
		defer wg.Done()
		shuffle.Byte(&input.CharSet)
	}()

	var index int
	var generated string
	pwdLen := input.Size - len(generated)

	wg.Wait()

	for i := 0; i < pwdLen; i++ {
		index, _ = strconv.Atoi(randomIndex[i])
		generated = generated + string(input.CharSet[index])
	}
	shuffle.String(&generated)

	return generated, credits, nil

}
