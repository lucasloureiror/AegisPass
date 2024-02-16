package shuffle

import "testing"

func TestBuildString(t *testing.T) {
	arr := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()")
	length := 16
	result := BuildString(arr, length)
	notExpected := "abcdefghijklmnop"
	if result == notExpected {
		t.Error("result showed the byte array in order into the string: ", result)
	}
}
