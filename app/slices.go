package goassessment

// SLICES
// For those functions that take a slice as input and return a slice,
// you can either modify the input slice or make a copy for the return
// the tests don't require that the input slice is not modified

// write a function that returns the index of an item in a slice
func indexOf(a []int, item int) int {
	for i, v := range a {
		// ff item is found, return the index
		if v == item {
			return i
		}
	}

	return -1
}

// write a function that sums the values in a slice
func sum(a []int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}

// write a function that removes all instances of a value from a slice
func remove(a []int, item int) []int {
	var result []int
	// iterate through the slice
	for _, v := range a {
		// add the element to result slice if it's not equal to the item
		if v != item {
			result = append(result, v)
		}
	}
	return result
}

// write a function that returns the value of the first element in a slice (wihtout removing it)
func front(a []int) int {
	return a[0]
}

// write a function that returns the value of the last element in a slice (wihtout removing it)
func back(a []int) int {
	return a[len(a)-1]
}

// write a function that adds an item to the end of a slice
func pushBack(a []int, item int) []int {
	return append(a, item)
}

// write a function that removes an item to the end of a slice
func popBack(a []int) []int {
	return a[:len(a)-1]
}

// write a function that adds an item to the front of a slice
func pushFront(a []int, item int) []int {
	return append([]int{item}, a...)
}

// write a function that removes an item from the front of a slice
func popFront(a []int) []int {
	return a[1:]
}

// write a function that concatenates two slices
func concat(a []int, b []int) []int {
	return append(a, b...)
}

// write a function that adds an item to a slice at the specified index
func insert(a []int, item int, index int) []int {
	a = append(a[:index], append([]int{item}, a[index:]...)...)
	return a
}

// write a function that returns a count of matching items in a slice
func count(a []int, item int) int {
	count := 0
	// loop through the slice and check if each element matches the item
	for _, v := range a {
		if v == item {
			count++
		}
	}
	return count
}

// write a function that finds duplicates in a slice
func duplicates(a []int) []int {
	seen := make(map[int]int) // map to store counts of elements
	var duplicates []int      // slice to store duplicates
	// loop through the slice and count occurrences
	for _, v := range a {
		seen[v]++
	}

	// find duplicates (elements that appear more than once)
	for key, count := range seen {
		if count > 1 {
			duplicates = append(duplicates, key)
		}
	}

	return duplicates
}

// write a function that squares all items in a slice
func square(a []int) []int {
	var result []int // slice to store the squared values
	// iterate over each element in the input slice
	for _, v := range a {
		result = append(result, v*v) // square the element and append to result
	}

	return result
}

// write a function that returns all the indices in a slice that matches an item
func findAllOccurrences(a []int, item int) []int {
	var indices []int // slice to store the indices of matching items
	// iterate through the slice
	for i, v := range a {
		if v == item { // check if the current item matches
			indices = append(indices, i) // add the index to the result slice
		}
	}

	return indices
}
