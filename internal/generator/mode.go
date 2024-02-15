package generator

import "github.com/lucasloureiror/AegisPass/internal/cli"

func ReturnGeneratorMode(data *cli.Input) PasswordGeneratorStrategy {

	if data.Flags.Offline {
		return offline{}
	}

	if data.Flags.UseStandard {
		return standard{}
	}

	return random{}

}
