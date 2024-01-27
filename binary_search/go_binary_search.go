//This code is for learning purposes.
//As a result, it may or may not be a little more heavily commented than normal.
package main

import (
	"fmt"
)

var orderedArray []int = []int{2, 4, 6, 8, 10, 12, 24, 36, 64, 128, 256, 512, 1024, 2048}

func binarySearch(array []int, value int)(presence bool, index int){
	var start int = 0
	var end int = len(array) - 1 
	var midpoint int //= end/2 <== I had this initial assignment here, but it's not needed when we have our 
			//updating assignment at the top of the loop instead of at the bottom.
	
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
						//If we start and end next to each other, and neither one of them is our value we could get stuck
						//in an infinite loop otherwise. Our midpoint would either be (start + 0) or (start + 1) which is
						//start and end respectively. It won't move and the start or end will just be reset over and over.
						//We need to let them keep moving so we can spot when they cross past each other. 
		}
	}
	return false, -1
}

func main() {
	g, h := binarySearch(orderedArray, 2)
	fmt.Println(g, h)
}
