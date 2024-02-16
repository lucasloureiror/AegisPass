/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package cli

type Input struct {
	Size              int
	Flags             Flags
	CharSet           []byte
	NumberOfPasswords int
}

type Flags struct {
	NeedHelp     bool
	UseOnlyNums  bool
	UseLower     bool
	UseUpper     bool
	UseSymbols   bool
	UseStandard  bool
	PrintCredits bool
	Online       bool
}
