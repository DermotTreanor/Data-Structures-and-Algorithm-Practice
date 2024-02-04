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
	var minimumItemIndex int 
	for currentStart:= 0; currentStart < len(array) - 1; currentStart++{
		minimumItemIndex = currentStart //This line is obviously necessary but I explain at the bottom what can go wrong without it

		for i := currentStart + 1; i < len(array); i++{ //Note that i is set to currentStart PLUS ONE. This starts i one ahead of currentStart (where minimum is initially set to)
														//This means we won't do an iteration just checking if the index minimum is already set to is less than itself.
			if array[i] < array[minimumItemIndex]{
				minimumItemIndex = i
			}
		}
		if minimumItemIndex != currentStart{
			array[currentStart], array[minimumItemIndex] = array[minimumItemIndex], array[currentStart]
		}
	}
}

func main(){
	var exampleArray []int = shuffle(orderedArray)
	fmt.Println(exampleArray)
	selectionSort(exampleArray)
	fmt.Println(exampleArray)
}



/*
Resetting the value that stores the index with the lowest value to be where we are currently starting our new main loop
iteration is very important. If we don't do that: it's not a big issue if the minimum is already ahead of the currentStart.
However, what if they are the same? As in, what if the minimum was where we started.
Then, we will increment currentStart but leave the minimum index where it is. This means minimum won't change in the inner
loop because it is part of the already sorted array (and therefore has a value lower than anything remaining).

But since our code only checks that we swap currentStart and the minimum if they aren't the same index we will be swapping
whatever happened to be the currentStart into minimum's place (the last minimum that we found will be placed at the new current).

Minimum will still be stuck where it and currentStart diverged. It will only leave if there are values in the array that are lower
than it so that it gets changed to that. 
*/