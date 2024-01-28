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

func bubble_sort(array []int){
	var unsortedLimit int = len(array) - 1
	var entryBubblingUp int 
	var entryToCompareTo int
	var stillSorting bool = true

	//The outer loop keeps going until we've had a pass through with no swaps
	for stillSorting{
		stillSorting = false
		entryBubblingUp = 0
		entryToCompareTo = entryBubblingUp + 1
		for entryToCompareTo <= unsortedLimit{
			fmt.Print("Test|")
			if array[entryBubblingUp] > array[entryToCompareTo]{
				stillSorting = true
				array[entryBubblingUp], array[entryToCompareTo] = array[entryToCompareTo], array[entryBubblingUp]
			}
			entryBubblingUp += 1
			entryToCompareTo += 1
		}
		unsortedLimit -= 1 //This is necessary to stop redundant checks of already sorted portions of the array
		fmt.Println("\nDURING: ", array)

	}
}

func main(){
	exampleArray := shuffle(orderedArray)
	fmt.Println("Here is the array BEFORE the sort: ", exampleArray)
	bubble_sort(exampleArray)
	fmt.Println("Here is the array AFTER the sort: ", exampleArray)
}

