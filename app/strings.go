package goassessment

import (
	"strings"
)

// write a function that composes a string from arguments
func composeString(a string, b string) string {
	return a + " " + b
}

// write a function that returns a string from a byte array
func fromBytes(b []byte) string {
	return string(b)
}

// write a function that splits a string into words
func splitString(s string) []string {
	words := strings.Fields(s)
	return words
}

// write a function that converts a string to Title Case
func titleString(s string) string {
	// split the string into words
	words := strings.Fields(s)
	for i, word := range words {
		// capitalize the first letter of each word
		words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
	}
	// join the words back together into a single string
	return strings.Join(words, " ")
}

func runesFromString(s string) []rune {
	return []rune(s)
}

// write a function that reduces adjacent repeated characters
func reduceString(s string, count int) string {
	if count <= 0 {
		return s // if count is 0 or negative, return the original string
	}

	var result strings.Builder
	var prevChar byte
	charCount := 0

	for i := 0; i < len(s); i++ {
		// if the character is the same as the previous one, increment the count
		if s[i] == prevChar {
			charCount++
		} else {
			// reset count for new character
			charCount = 1
			prevChar = s[i]
		}

		// if count is less than or equal to the limit, append the character
		if charCount <= count {
			result.WriteByte(s[i])
		}
	}

	return result.String()
}

// write a function that wraps lines at a given number of columns without breaking works
func wordWrap(s string, column int) string {
	var result []string
	var currentLine []string
	words := strings.Fields(s) // split the string into words

	for _, word := range words {
		// check if adding this word would exceed the column limit
		lineLength := len(strings.Join(currentLine, " ")) + len(word)
		if lineLength <= column {
			// add the word to the current line
			currentLine = append(currentLine, word)
		} else {
			// if line exceeds column limit, push current line to result
			result = append(result, strings.Join(currentLine, " "))
			// start a new line with the current word
			currentLine = []string{word}
		}
	}

	// add the last line
	if len(currentLine) > 0 {
		result = append(result, strings.Join(currentLine, " "))
	}

	// join all lines into a final result with newlines
	return strings.Join(result, "\n")
}

// write a function that reverses the characters in a string
func reverseString(s string) string {
	// convert the string to a slice of runes to handle multi-byte characters correctly
	runes := []rune(s)

	// reverse the slice of runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
