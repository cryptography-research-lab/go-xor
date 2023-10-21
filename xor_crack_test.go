package xor

import (
	"testing"
)

func TestCrack(t *testing.T) {
	key := []byte("CC11001100")

	plainSliceA := []byte("密码学")
	cipherSliceA := XOR(plainSliceA, key)

	plainSliceB := []byte("真有趣")
	cipherSliceB := XOR(plainSliceB, key)

	crackResultKey := Crack(plainSliceA, cipherSliceA, cipherSliceB)
	t.Log(string(crackResultKey))

}
