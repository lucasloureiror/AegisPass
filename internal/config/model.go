/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package config

type Password struct {
	Size      int
	Flags     Flags
	APICredit int
	Generated string
	CharSet   []byte
}

type Flags struct {
	NeedHelp     bool
	UseOnlyNums  bool
	UseLower     bool
	UseUpper     bool
	UseSymbols   bool
	UseStandard  bool
	PrintCredits bool
	Offline      bool
}
