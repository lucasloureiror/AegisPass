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
	OfflineMode  bool
}
