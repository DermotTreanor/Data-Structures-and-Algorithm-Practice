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
	//We start at the second position and place the item into the correct, sorted spot with everything before it.
	//As we move through the list, each index we get to will have everything to the 'left' of it sorted so we just need to
	//find its own correct place
	var currentStartIndex int = 1
	var currentValue int
	var gapIndex int

	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~This was the original attempt
	// for currentStartIndex < len(array) {
	// 	gapIndex = currentStartIndex
	// 	currentValue = array[gapIndex]
	// 	for comparisonIndex := gapIndex - 1; comparisonIndex >= -1; comparisonIndex-- {
	// 		if comparisonIndex == -1 {
	// 			array[gapIndex] = currentValue
	// 		} else if array[comparisonIndex] > currentValue {
	// 			array[gapIndex] = array[comparisonIndex]
	// 			gapIndex = comparisonIndex
	// 		} else {
	// 			array[gapIndex] = currentValue
	// 			break
	// 		}
	// 	}
	// 	currentStartIndex++
	// }


	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~We realised we only need to fill the gap in the outer loop
	// for currentStartIndex < len(array) {
	// 	gapIndex = currentStartIndex
	// 	currentValue = array[gapIndex]
	// 	for comparisonIndex := gapIndex - 1; comparisonIndex >= 0; comparisonIndex-- {
	// 		if array[comparisonIndex] > currentValue {
	// 			array[gapIndex] = array[comparisonIndex]
	// 			gapIndex = comparisonIndex
	// 		} else {
	// 			break
	// 		}
	// 	}
	// 	array[gapIndex] = currentValue
	// 	currentStartIndex++
	// }


	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~The loop is built around gapIndex
	// for currentStartIndex < len(array) {
	// 	gapIndex = currentStartIndex
	// 	currentValue = array[gapIndex]
	// 	for gapIndex >= 1{
	// 		comparisonIndex := gapIndex - 1
	// 		if array[comparisonIndex] > currentValue {
	// 			array[gapIndex] = array[comparisonIndex]
	// 			gapIndex--
	// 		} else {
	// 			break
	// 		}
	// 	}
	// 	array[gapIndex] = currentValue  //Taking this out of the nested loop is the key.
	// 	currentStartIndex++
	// }

	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~We can remove comparisonIndex completely and just 
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~reference that index as 'gapIndex - 1' (it's only needed twice)
	for currentStartIndex < len(array) {
		gapIndex = currentStartIndex
		currentValue = array[gapIndex]
		for gapIndex >= 1{
			if array[gapIndex - 1] > currentValue {
				array[gapIndex] = array[gapIndex - 1]
				gapIndex--
			} else {
				break
			}
		}
		array[gapIndex] = currentValue  //Taking this out of the nested loop is the key.
		currentStartIndex++
	}



}

func main() {
	exampleArray = shuffle(exampleArray)
	fmt.Println("Here is the shuffled array: ", exampleArray)
	InsertionSort(exampleArray)
	fmt.Println("This is the new, sorted array: ", exampleArray)
}
