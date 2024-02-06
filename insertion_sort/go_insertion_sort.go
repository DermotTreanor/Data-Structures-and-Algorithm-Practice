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

	var gapIndex int = 1
	var currentValue int
	for gapIndex < len(array)-1 {
		currentValue = array[gapIndex]
		for i := gapIndex - 1; i >= 0; i-- {
			if array[i] > array[gapIndex]{
				array[gapIndex] = array[i]
				gapIndex = i
			}else{
				array[gapIndex] = currentValue
				break
			}
		}
		gapIndex++

	}

}

func main() {
	// fmt.Println("Here is the ordered array: ", exampleArray)
	exampleArray = shuffle(exampleArray)
	fmt.Println("Here is the shuffled array: ", exampleArray)
	InsertionSort(exampleArray)
	fmt.Println("This is the new, sorted array: ", exampleArray)
}
