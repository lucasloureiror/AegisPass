package generator

import "github.com/lucasloureiror/AegisPass/internal/cli"

func ReturnGeneratorMode(data *cli.Input) PasswordGeneratorStrategy {

	if data.Flags.NeedHelp {
		return help{}
	}

	if data.Flags.UseStandard {
		return standard{}
	}

	if data.Flags.Online {
		return online{}
	}

	return random{}

}
