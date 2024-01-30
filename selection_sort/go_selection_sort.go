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

	var minimumItemIndex int = 0

	for i, v := range array{
		if v < array[minimumItemIndex] && i != minimumItemIndex{
			minimumItemIndex = i
		}
	}
		

}

func main(){
	var exampleArray []int = shuffle(orderedArray)
	fmt.Println(exampleArray)
	selectionSort(exampleArray)
	fmt.Println(exampleArray)

}