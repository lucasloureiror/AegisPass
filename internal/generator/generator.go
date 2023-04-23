package generator

import (
	"github.com/lucasloureiror/AegisPass/internal/charsets"
	"github.com/lucasloureiror/AegisPass/internal/config"
	"github.com/lucasloureiror/AegisPass/internal/output"
	"github.com/lucasloureiror/AegisPass/internal/randomclient"
	"github.com/lucasloureiror/AegisPass/internal/shuffle"
)

type generator interface {
	generate(*config.Password, []string)
}

func Start(pwd *config.Password) {
	apiResponse := make(chan []string)

	go func() {
		apiResponse <- randomclient.Init(pwd)
	}()

	charsets.Create(pwd)

	shuffle.Byte(&pwd.CharSet)

	var gen generator
	if pwd.Flags.UseStandard {
		gen = &standard{}
		gen.generate(pwd, <-apiResponse)
	} else {
		gen = &random{}
		gen.generate(pwd, <-apiResponse)
	}

	output.Print(pwd)

}
