package xor

import "github.com/golang-infrastructure/go-gtypes"

// Crack 根据A的明文和密文，以及B的密文，计算出B的明文，密文A和密文B需要是用一个秘钥加密得出的
func Crack[T gtypes.Unsigned](plainSliceA, cipherSliceA, cipherSliceB []T) []T {
	r := XOR[T](cipherSliceA, cipherSliceB)
	return XOR[T](r, plainSliceA)
}
