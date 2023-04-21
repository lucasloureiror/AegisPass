package generator

import (
	"strconv"

	"github.com/lucasloureiror/AegisPass/internal/charsets"
	"github.com/lucasloureiror/AegisPass/internal/config"
	"github.com/lucasloureiror/AegisPass/internal/output"
	"github.com/lucasloureiror/AegisPass/internal/randomclient"
	"github.com/lucasloureiror/AegisPass/internal/shuffle"
)

func Start(pwd *config.Password) {

	apiResponse := make(chan []string)

	go func() {
		apiResponse <- randomclient.Init(pwd)
	}()

	charsets.Create(pwd)

	shuffle.Byte(&pwd.CharSet)

	if pwd.Flags.UseStandard {
		makeStdRandomPass(pwd)
	}
	makeRandomPass(pwd, <-apiResponse)

	output.Print(pwd)

}

func makeRandomPass(pwd *config.Password, randomInt []string) {

	var index int
	pwdLen := pwd.Size - len(pwd.Generated)

	for i := 0; i < pwdLen; i++ {
		index, _ = strconv.Atoi(randomInt[i])
		pwd.Generated = pwd.Generated + string(pwd.CharSet[index])
	}
	shuffle.String(&pwd.Generated)
}

func makeStdRandomPass(pwd *config.Password) {

	upper := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	shuffle.Byte(&upper)
	special := []byte("!@#$%&*")
	shuffle.Byte(&special)
	nums := []byte("0123456789")
	shuffle.Byte(&nums)

	pwd.Generated = string(upper[0]) + string(special[0]) + string(nums[0])

}
