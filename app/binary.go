package goassessment

import (
	"fmt"
	"strconv"
)

// write a function that returns the 0 or 1 value at the specified bit position
func valueAtbit(num int, bit int) int {
	return (num >> bit) & 1
}

// write a function that returns the base 10 value of the binary string
func base10(n string) int {
	// parse the string as a binary number
	result, err := strconv.ParseInt(n, 2, 0)
	if err != nil {
		// handle error if the string is not a valid binary number
		fmt.Println("Error:", err)
		return -1
	}
	return int(result)
}

// write a function that returns the binary value of num as a string
func convertToBinary(num int) string {
	var binary string
	for num > 0 {
		binary = fmt.Sprintf("%d", num%2) + binary
		num = num / 2
	}
	return binary
}

// write a function that returns the bitwise OR of the input values
func bitwiseOr(x int, y int, z int) int {
	return x | y | z
}

// write a function that returns the bitwise AND of the input values
func bitwiseAnd(x int, y int, z int) int {
	return x & y & z
}

// write a function that returns the bitwise XOR of the input values
func bitwiseXor(x int, y int, z int) int {
	return x ^ y ^ z
}
