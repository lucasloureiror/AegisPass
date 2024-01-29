package generator

import (
	"github.com/lucasloureiror/AegisPass/internal/charsets"
	"github.com/lucasloureiror/AegisPass/internal/config"
	"github.com/lucasloureiror/AegisPass/internal/output"
	"github.com/lucasloureiror/AegisPass/internal/randomclient"
	"github.com/lucasloureiror/AegisPass/internal/shuffle"
)

type generator interface {
	generate(*config.Password, []string) //Passing the model for a password ande slice of strings with api generated numbers.
}

func Start(pwd *config.Password) {

	var gen generator
	charsets.Create(pwd)

	if pwd.Flags.Offline {
		gen = offline{}
		gen.generate(pwd, nil)
		output.Print(pwd)
		return
	}

	apiResponse := make(chan []string)

	go func() {
		apiResponse <- randomclient.Start(pwd)
	}()

	shuffle.Byte(&pwd.CharSet)

	if pwd.Flags.UseStandard {
		gen = standard{}
		gen.generate(pwd, <-apiResponse)
	} else {
		gen = random{}
		gen.generate(pwd, <-apiResponse)
	}

	output.Print(pwd)

}
