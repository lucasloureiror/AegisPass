package shuffle

func Byte(set *[]byte) {

	arrSize := len(*set) - 1

	for i := arrSize; i > 0; i-- {
		j := randomInt(arrSize)
		(*set)[i], (*set)[j] = (*set)[j], (*set)[i]
	}

}
