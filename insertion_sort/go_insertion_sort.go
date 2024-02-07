package main

import (
	"fmt"
	"math/rand"
	"time"
)

var exampleArray []int = []int{2, 4, 6, 8, 10, 12, 24, 36, 64, 128, 512, 1024, 2048}

func shuffle(array []int) (shuffledArray []int) {
	shuffledArray = make([]int, len(array))
	randomGenerator := rand.New(rand.NewSource(time.Now().Unix()))
	shuffledIndices := randomGenerator.Perm(len(array))
	for i := 0; i < len(array); i++ {
		shuffledArray[i] = array[shuffledIndices[i]]
	}

	return shuffledArray
}

func InsertionSort(array []int) {
	var currentStartIndex int = 1
	var currentValue int
	var gapIndex int

	//We start at the second position and place the item into the correct, sorted spot with everything before it.
	//As we move through the list, each index we get to will have everything to the 'left' of it sorted so we just need to
	//find its own correct place.
	for currentStartIndex < len(array) {
		gapIndex = currentStartIndex
		currentValue = array[gapIndex]
		for comparisonIndex := gapIndex - 1; comparisonIndex >= -1; comparisonIndex-- {
			if comparisonIndex == -1{
				array[gapIndex] = currentValue
			}else if array[comparisonIndex] > currentValue{
				array[gapIndex] = array[comparisonIndex]
				gapIndex = comparisonIndex
			}else{
				array[gapIndex] = currentValue
				break
			}
		}
		currentStartIndex++
	}
}

func main() {
	exampleArray = shuffle(exampleArray)
	fmt.Println("Here is the shuffled array: ", exampleArray)
	InsertionSort(exampleArray)
	fmt.Println("This is the new, sorted array: ", exampleArray)
}
