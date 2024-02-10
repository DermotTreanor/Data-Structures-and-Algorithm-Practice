package main

import (
	"fmt"
)

var orderedArray []int = []int{2, 4, 6, 8, 10, 12, 24, 36, 64, 128, 256, 512, 1024, 2048}

func binarySearch(array []int, value int)(presence bool, index int){
	var start int = 0
	var end int = len(array) - 1 
	var midpoint int 

	//It's important we are only checking when the start and end indices have move passed each other.
	//If we just check when they are equal we might have the indices land on each other after one midpoint value
	//is found to be wrong and then, before we've checked the value they are both on, break out of the loop too early.
	for end >= start{
		midpoint = start + (end-start)/2
		if array[midpoint] == value{
			return true, midpoint
		} else if array[midpoint] > value{
			end = midpoint - 1
		} else if array[midpoint] < value{ 
			start = midpoint + 1	//Our start and end positions are always set just after or before the midpoint respectively. 
		}
	}
	return false, -1
}

func main() {
	inArray, index := binarySearch(orderedArray, 25)
	fmt.Println(inArray, index)
}
