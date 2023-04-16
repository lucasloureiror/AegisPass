package shuffle

import (
	"crypto/rand"
	"math/big"
)

func ShuffleByte(set *[]byte) {

	arrSize := len(*set)

	for i := arrSize - 1; i > 0; i-- {
		jBigInt, err := rand.Int(rand.Reader, big.NewInt(int64(arrSize)))
		if err != nil {
			panic(err)
		}
		j := int(jBigInt.Int64())
		(*set)[i], (*set)[j] = (*set)[j], (*set)[i]
	}

}

func ShuffleStr(pwd *string) {

	pwdBytes := []byte(*pwd)
	strSize := len(pwdBytes)

	for i := strSize - 1; i > 0; i-- {
		jBigInt, err := rand.Int(rand.Reader, big.NewInt(int64(strSize)))
		if err != nil {
			panic(err)
		}
		j := int(jBigInt.Int64())
		pwdBytes[i], pwdBytes[j] = pwdBytes[j], pwdBytes[i]
	}

	*pwd = string(pwdBytes)

}
