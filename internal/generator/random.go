package generator

import (
	"strconv"

	"github.com/lucasloureiror/AegisPass/internal/config"
	"github.com/lucasloureiror/AegisPass/internal/shuffle"
)

type random struct{}

func (random) generate(pwd *config.Password, randomIndex []string) {

	var index int
	pwdLen := pwd.Size - len(pwd.Generated)

	for i := 0; i < pwdLen; i++ {
		index, _ = strconv.Atoi(randomIndex[i])
		pwd.Generated = pwd.Generated + string(pwd.CharSet[index])
	}
	shuffle.String(&pwd.Generated)

}
