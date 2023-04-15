package generator

import (
	"strconv"

	"github.com/lucasloureiror/AegisPass/internal/charsets"
	"github.com/lucasloureiror/AegisPass/internal/config"
	"github.com/lucasloureiror/AegisPass/internal/output"
	"github.com/lucasloureiror/AegisPass/internal/randomclient"
	"github.com/lucasloureiror/AegisPass/internal/shuffle"
)

func Init(pwd *config.Password) {

	apiResponse := make(chan []string)

	go func() {
		apiResponse <- randomclient.Init(pwd)
	}()

	charsets.Create(pwd)

	shuffle.Shuffle(pwd)

	makeRandomPass(pwd, <-apiResponse)

	output.Print(pwd)

}

func makeRandomPass(pwd *config.Password, randomInt []string) {

	var index int

	for i := 0; i < pwd.Size; i++ {
		index, _ = strconv.Atoi(randomInt[i])
		pwd.Generated = pwd.Generated + string(pwd.CharSet[index])
	}

}
