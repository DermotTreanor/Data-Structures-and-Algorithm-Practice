package main

import (
	"fmt"
	"math/rand"
	"time"
)

func get_mixed_array() (mixed []int) {
	ordered_array := []int{2, 4, 6, 8, 10, 12, 24, 36, 64, 128, 512, 1024, 2048}
	fmt.Println("Here is unshuffled: ", ordered_array)
	mixed = make([]int, len(ordered_array))

	random_gen := rand.New(rand.NewSource(time.Now().Unix()))
	rand_perm := random_gen.Perm(len(ordered_array))
	for i := 0; i < len(rand_perm); i++ {
		mixed[i] = ordered_array[rand_perm[i]]
	}
	fmt.Println("Here is SHUFFLED: ", mixed)
	fmt.Print("\n\n")
	return mixed
}

type PartitionableArray struct {
	array []int
}


/*
Note: In this method we are passing a pointer of our object. This isn't actually important here since the underlying
structure is a slice. 
If this were an array then we would HAVE to do this because passing in a copy of our object would change a copy of the array.
As it stands, whether we pass in a copy of our object only affects whether we are copying the reference of a slice.
But the reference will still point to the same data whether it is a copy or not.
*/
func (p_array *PartitionableArray) Partition(lower_ind int, upper_ind int) {
	if (upper_ind - lower_ind) < 1{
		return
	}

	original_upper := upper_ind
	original_lower := lower_ind
	pivot_index := upper_ind
	pivot_value := p_array.array[pivot_index]
	upper_ind -= 1
	for{
		for lower_ind < len(p_array.array) - 1 && !(p_array.array[lower_ind] >= pivot_value){
			lower_ind++
		}
		for upper_ind >= 0 && !(p_array.array[upper_ind] <= pivot_value){
			upper_ind--
		}
		if lower_ind >= upper_ind{
			p_array.array[lower_ind], p_array.array[pivot_index] = p_array.array[pivot_index], p_array.array[lower_ind]
			break
		} else{
			p_array.array[lower_ind], p_array.array[upper_ind] = p_array.array[upper_ind], p_array.array[lower_ind]
		}
	}

	p_array.Partition(original_lower, lower_ind - 1)
	p_array.Partition(lower_ind + 1, original_upper)

}

func (p_array *PartitionableArray)Quicksort(){
	p_array.Partition(0, len(p_array.array))
	return
}


func main() {
	test_array := PartitionableArray{
		array: get_mixed_array(),
	}
	fmt.Println(test_array.array)
	test_array.Partition(0, len(test_array.array) - 1)
	fmt.Println(test_array.array)

	return
}
