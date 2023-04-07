package config

type Password struct {
	Size        int
	UseOnlyNums bool
	UseLower    bool
	UseUpper    bool
	UseSymbols  bool
	UseStandard bool
	Offline     bool
	Generated   string
	CharSet     []byte
}
