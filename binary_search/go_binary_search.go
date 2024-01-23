package main

import (
	"fmt"
	"math"
)

var orderedArray []int = []int{2, 4, 6, 8, 10, 12, 24, 36, 64, 128, 256, 512, 1024, 2048}

func binarySearch(array []int, value int)(presence bool, index int){
	var start int = 0
	var end int = len(array)

	// var midpoint int = end/2	//We always round down.
	//We round appropriately so we are always either in the middle (rounding up) or just to the left of middle (when it's even)
	var midpoint int = int(math.Round(float64(end)/2))
	for end != start{
		if array[midpoint] == value{
			return true, midpoint
		} else if array[midpoint] > value{
			end = midpoint
			midpoint = end/2
		} else if array[midpoint] < value{
			start = midpoint
			midpoint = (end-start)/2
		}
	}
	fmt.Println(midpoint)
	return false, 0
}

func main() {
	g, h := binarySearch(orderedArray, 10)
	fmt.Println(g, h)
}
