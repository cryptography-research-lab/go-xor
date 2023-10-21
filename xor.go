package xor

import (
	"github.com/golang-infrastructure/go-gtypes"
	variable_parameter "github.com/golang-infrastructure/go-variable-parameter"
)

// Options xor时使用的选项
type Options struct {

	// 当XOR的两个参数的长度不相等的话，如果需要交换的话，则以长的为准交换到第一个，否则默认是以第一个参数为准不考虑第二个参数的长度
	SwapForSort bool
}

func XOR[T gtypes.Unsigned](sliceA, sliceB []T, options ...*Options) []T {

	// 设置默认参数
	variable_parameter.SetDefaultParamByFunc(options, func() *Options {
		return &Options{
			SwapForSort: false,
		}
	})

	if options[0].SwapForSort {
		// 保证B是长度较短的那个
		if len(sliceB) > len(sliceA) {
			sliceA, sliceB = sliceB, sliceA
		}
	}

	resultSlice := make([]T, 0)
	for index, valueA := range sliceA {
		valueB := sliceB[index%len(sliceB)]
		resultSlice = append(resultSlice, valueA^valueB)
	}
	return resultSlice
}

// Encrypt XOR加密
func Encrypt[T gtypes.Unsigned](plainSlice, keySlice []T) []T {
	return XOR[T](plainSlice, keySlice)
}

// Decrypt XOR解密
func Decrypt[T gtypes.Unsigned](cipherSlice, keySlice []T) []T {
	return XOR[T](cipherSlice, keySlice)
}
