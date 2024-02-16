package shuffle

import "testing"

func TestRandomInt(t *testing.T) {
	max := 10
	result := randomInt(max)
	if result > max {
		t.Error("result is greater than max: ", result)
	}
}
