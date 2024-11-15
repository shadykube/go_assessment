package goassessment

import (
	"fmt"
	"strings"
)

// write a function that returns a function
func fFunction(str string) func(string) string {
	return func(suffix string) string {
		firstWord := str
		newString := firstWord + ", " + suffix
		return newString
	}
}

// write a function that returns a slice of closures
func fMakeClosures(fn func(x int) int, arr []int) []func() int {
	closures := make([]func() int, len(arr))
	for i, val := range arr {
		// Create a closure that captures `val`
		closures[i] = func(v int) func() int {
			return func() int {
				return fn(v)
			}
		}(val)
	}
	return closures
}

// write a function that implements a partial application
func fPartial(
	fn func(a string, b string, c string) string,
	str1 string,
	str2 string,
) func(string) string {
	return func(str3 string) string {
		return fn(str1, str2, str3)
	}
}

// write a function that creates a new array populated
// with the result of calling the provided function
func fMap(fn func(a int) int, arr []int) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = fn(v)
	}
	return result
}

// write a function that applies the reducer to a slice of integers
func fReduce(fn func(acc int, val int) int, arr []int) int {
	acc := arr[0]
	for _, val := range arr[1:] {
		acc = fn(acc, val)
	}
	return acc
}

// write a function that applies the filter to a slice of integers
func fFilter(fn func(a int) bool, arr []int) []int {
	var result []int
	for _, val := range arr {
		if fn(val) {
			result = append(result, val)
		}
	}
	return result
}

// write a function that handles a variadic argument
func fVariadic(s ...int) string {
	var strNumbers []string
	for _, num := range s {
		strNumbers = append(strNumbers, fmt.Sprintf("%d", num))
	}
	return strings.Join(strNumbers, ",")
}
