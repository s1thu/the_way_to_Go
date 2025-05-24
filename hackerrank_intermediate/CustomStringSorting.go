package main

import (
	"fmt"
	"sort"
)

// 3. Go: Custom String Sorting
// election has been disabled in this region
// Sort a set of strings based on the following factors:
// 1. An odd length string should precede an even length string.
// 2. If both strings have odd lengths, the shorter of the two should precede.
// 3. If both strings have even lengths, the longer of the two should precede.
// 4. If the two strings have equal lengths, they should be in alphabetical order.
// Example
// strArr =['abc', 'ab', 'abcde', 'a', 'abcd']
// The result is ['a', 'abc', 'abcde', 'abcd', 'ab']. Odd length words are sorted in ascending order of length followed by even length words in descending order of length. There are no matching length words, so no alphabetical sorting is required.
func customSort(str []string) []string {
	oddSortedArray := make([]string, 0)
	evenSortedArray := make([]string, 0)

	for _, s := range str {
		if len(s)%2 != 0 {
			oddSortedArray = append(oddSortedArray, s)
		} else {
			evenSortedArray = append(evenSortedArray, s)
		}
	}

	// Sort odd by ascending length; if equal, alphabetical
	sort.Slice(oddSortedArray, func(i, j int) bool {
		if len(oddSortedArray[i]) == len(oddSortedArray[j]) {
			return oddSortedArray[i] < oddSortedArray[j]
		}
		return len(oddSortedArray[i]) < len(oddSortedArray[j])
	})

	// Sort even by descending length; if equal, alphabetical
	sort.Slice(evenSortedArray, func(i, j int) bool {
		if len(evenSortedArray[i]) == len(evenSortedArray[j]) {
			return evenSortedArray[i] < evenSortedArray[j]
		}
		return len(evenSortedArray[i]) > len(evenSortedArray[j])
	})

	// Merge
	result := append(oddSortedArray, evenSortedArray...)
	return result
}

func main() {
	strArr := []string{"abc", "ab", "abcde", "a", "abcd"}
	result := customSort(strArr)
	fmt.Println(result) // Output: [a abc abcde abcd ab]
}
