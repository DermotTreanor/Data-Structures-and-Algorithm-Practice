package main

import (
	"fmt"
)

func anagrams(word string) int {
	if N == 0 {
		return 1
	} else if N < 0 {
		return 0
	}
	return staircaseRecursive(N-1) + staircaseRecursive(N-2) + staircaseRecursive(N-3)
}

func main() {
	fmt.Println(anagrams("abc"))
}
