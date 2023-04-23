package generator

import (
	"github.com/lucasloureiror/AegisPass/internal/config"
	"github.com/lucasloureiror/AegisPass/internal/shuffle"
)

type offline struct{}

func (offline) generate(pwd *config.Password, _ []string) {
	shuffle.Byte(&pwd.CharSet)
	pwdByte := shuffle.FisherYates(pwd.CharSet, pwd.Size)
	pwd.Generated = string(pwdByte)
}
