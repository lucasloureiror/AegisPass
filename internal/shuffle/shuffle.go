package shuffle

import (
	"crypto/rand"
	"math/big"
)

func Shuffle() []byte {

	charSet := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*")
	arrSize := len(charSet)

	for i := arrSize - 1; i > 0; i-- {
		jBigInt, err := rand.Int(rand.Reader, big.NewInt(int64(arrSize)))
		if err != nil {
			panic(err)
		}
		j := int(jBigInt.Int64())
		charSet[i], charSet[j] = charSet[j], charSet[i]
	}

	return charSet

}
