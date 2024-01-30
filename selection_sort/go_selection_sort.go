package main

import (
	"fmt"
	"math/rand"
	"time"
)

var orderedArray []int = []int{2, 4, 6, 8, 10, 12, 24, 36, 64, 128, 256, 512, 1024, 2048}
func shuffle(array []int)(randSlice []int){
	randSlice = make([]int, len(array))
	randomGenerator := rand.New(rand.NewSource(time.Now().Unix()))
	var randomIndices []int = randomGenerator.Perm(len(array))
	for i, v := range randomIndices{
		randSlice[i] = array[v]
	}
	return randSlice
}

func selectionSort(array []int){
	/*
	In selection sort we run through an array and keep track of the lowest value's index 
	(updating it if we find a value that's lower).
	At the end of the run we will swap that index's value to the beginning of the array. 
	We then run through again, but this time we start from the second position (since the first index now has the minimum).
	We end up placing the second lowest item there. This pattern is continued until are starting position is the last item
	in the array.
	*/
	var minimumItemIndex int = 0
	var currentStart int = 0
	for currentStart != len(array) - 1{
		for i := currentStart; i < len(array); i++{
			if array[i] < array[minimumItemIndex]{
				minimumItemIndex = i
			}
		}
		array[currentStart], array[minimumItemIndex] = array[minimumItemIndex], array[currentStart]
		currentStart++
	}
}

func main(){
	var exampleArray []int = shuffle(orderedArray)
	fmt.Println(exampleArray)
	selectionSort(exampleArray)
	fmt.Println(exampleArray)
}