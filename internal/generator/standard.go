package generator

import (
	"strconv"

	"github.com/lucasloureiror/AegisPass/internal/config"
	"github.com/lucasloureiror/AegisPass/internal/shuffle"
)

type standard struct{}

func (s *standard) generate(pwd *config.Password, randomIndex []string) {
	upper := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	shuffle.Byte(&upper)
	special := []byte("!@#$%&*")
	shuffle.Byte(&special)
	nums := []byte("0123456789")
	shuffle.Byte(&nums)

	pwd.Generated = string(upper[0]) + string(special[0]) + string(nums[0])

	var index int
	pwdLen := pwd.Size - len(pwd.Generated)

	for i := 0; i < pwdLen; i++ {
		index, _ = strconv.Atoi(randomIndex[i])
		pwd.Generated = pwd.Generated + string(pwd.CharSet[index])
	}
	shuffle.String(&pwd.Generated)
}
