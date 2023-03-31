package shuffle

import(
	"crypto/rand"
	"math/big"
)

func Shuffle()([]byte){

	allChars := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*")
	arraySize := len(allChars)

	for i := arraySize - 1; i > 0; i-- {
		jBigInt, err := rand.Int(rand.Reader, big.NewInt(int64(arraySize)))
		if err != nil {
			panic(err)
		}
		j := int(jBigInt.Int64())
		allChars[i], allChars[j] = allChars[j], allChars[i]
	}

	return allChars

}
