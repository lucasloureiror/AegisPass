package shuffle

import (
	"crypto/rand"
	"github.com/lucasloureiror/AegisPass/internal/config"
	"math/big"
)

func Shuffle(pwd *config.Password) {

	arrSize := len(pwd.CharSet)

	for i := arrSize - 1; i > 0; i-- {
		jBigInt, err := rand.Int(rand.Reader, big.NewInt(int64(arrSize)))
		if err != nil {
			panic(err)
		}
		j := int(jBigInt.Int64())
		pwd.CharSet[i], pwd.CharSet[j] = pwd.CharSet[j], pwd.CharSet[i]
	}

}
