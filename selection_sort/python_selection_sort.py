#! /usr/bin/python3

import random
my_list = [2, 4, 6, 8, 10, 12, 24, 36, 64, 128, 256, 512, 1024, 2048]
random.shuffle(my_list)

def selection_sort(array):
    currentStart = 0

    while currentStart < len(array) - 1:
        minimumItemIndex = currentStart #Interesting glitch if this isn't here as currentStart will be ahead of minimum for the next run. 
                                        #Minimum will be guaranteed not to swap with anything on this iteration. But it won't be == to currentStart
                                        #so we will swap what's at minimum back into the pool of items to be sorted and take whatever currentStart is at
                                        #into our previously sorted pool. 
        for i, v in enumerate(array[currentStart + 1:]): #We start from currentStart PLUS ONE because minimumItemIndex is already set to currentStart. No point checking for a lower value at currentStart
            i += currentStart + 1
            if v < array[minimumItemIndex]:
                minimumItemIndex = i
        if minimumItemIndex != currentStart:
            array[currentStart], array[minimumItemIndex] = array[minimumItemIndex], array[currentStart]
        currentStart += 1
    
    return

def main():
    print("Here is the shuffled list: ", my_list)
    selection_sort(my_list)
    print("Here is the bubble sorted list: ", my_list)
    return 

if __name__ == "__main__":
    main()