package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func create_random_array(quantity int) (random_array []int) {
	ordered_array := []int{2, 4, 6, 8, 10, 12, 24, 36, 64, 128, 512, 1024, 2048}
	random_array = make([]int, len(ordered_array))
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	jumbled_indices := randomizer.Perm(len(random_array))
	for i, v := range jumbled_indices {
		random_array[v] = ordered_array[i]
	}

	return random_array
	//return []int{128, 4, 64, 6, 512, 24, 2048, 1024, 8, 12, 10, 36, 2}
}

func Partition(array []int, l_point int, r_point int) (partition_index int) {
	pivot_index := r_point
	r_point = r_point - 1
	for {
		for array[l_point] < array[pivot_index] {
			l_point++
		}
		//This loop needs to have a check so that r_point doesn't go past the beginning.
		//In another language it might stop because when r_point is -1 we refer to the pivot_index.
		//But Go doesn't allow for negative indices.
		for array[r_point] > array[pivot_index] && r_point > 0 {
			r_point--
		}
		if l_point >= r_point {
			break
		} else {
			array[l_point], array[r_point] = array[r_point], array[l_point]
		}
	}
	array[l_point], array[pivot_index] = array[pivot_index], array[l_point]
	partition_index = l_point
	return partition_index
}

func Quickselect(array []int, Nth_lowest int, lower_bound int, upper_bound int) (value int) {
	partition_index := Partition(array, lower_bound, upper_bound)
	if Nth_lowest-1 == partition_index {
		return array[partition_index]
	} else if Nth_lowest-1 < partition_index {
		return Quickselect(array, Nth_lowest, lower_bound, partition_index-1)
	} else {
		return Quickselect(array, Nth_lowest, partition_index+1, upper_bound)
	}

}

var nth = flag.Int("n", 4, "Select the nth lowest value to find")

func main() {
	flag.Parse()
	array_to_select := create_random_array(10)
	fmt.Println(array_to_select)
	nth_result := Quickselect(array_to_select, *nth, 0, len(array_to_select)-1)
	fmt.Println(array_to_select)
	fmt.Printf("The fourth lowest is: %d\n", nth_result)

}
