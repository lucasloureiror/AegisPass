package shuffle

import (
	"crypto/rand"
	"encoding/binary"
)

// Well, when I came back to this code the logic was already in place and didn't want to change.
func fisherYatesSelector(arr []byte, length int) []byte {
	shuffled := make([]byte, length)
	for i := 0; i < length; i++ {
		j := randomInt(len(arr))
		shuffled[i] = arr[j]
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
