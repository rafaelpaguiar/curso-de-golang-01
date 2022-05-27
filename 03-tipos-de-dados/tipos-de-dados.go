package main

import (
	"errors"
	"fmt"
)

func main() {

	var int1 int64 = -100000000000000000
	var uint2 uint32 = 10000

	var rune1 = 12345

	var byte1 = 123

	fmt.Println(int1, uint2, rune1, byte1, "\n")

	var float1 float32 = 0.32
	var float2 float64 = 0.64

	fmt.Println(float1, float2, "\n")

	var str string = "String"

	fmt.Println(str, "\n")

	char := 'A'

	fmt.Println(char, "\n")

	var stringVazia string

	fmt.Println(stringVazia, "\n")

	var intVazio int16
	fmt.Println(intVazio, "\n")

	var bool1 bool = true
	fmt.Println(bool1, "\n")

	var err error = errors.New("Erro declarado")
	fmt.Println(err, "\n")

}
