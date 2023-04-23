package shuffle

import (
	"crypto/rand"
	"encoding/binary"
)

func fisherYatesShuffle(arr []byte, length int) []byte {
	shuffled := make([]byte, length)
	copy(shuffled, arr)

	for i := 0; i < length; i++ {
		j := randomInt(len(arr))
		shuffled[i] = arr[j]
		arr = append(arr[:j], arr[j+1:]...)
	}

	return shuffled
}

func randomInt(max int) int {
	buf := make([]byte, 4)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	val := int(binary.BigEndian.Uint32(buf))
	return val % max
}
