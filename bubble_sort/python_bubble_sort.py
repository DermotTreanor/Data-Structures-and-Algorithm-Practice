#! /usr/bin/python3
import random
my_list = [2, 4, 6, 8, 10, 12, 24, 36, 64, 128, 256, 512, 1024, 2048]
random.shuffle(my_list)


def bubble_sort(array):
    endpoint = len(array)
    haveSwaps = True
    while haveSwaps == True:
        left_point = 0
        right_point = left_point + 1
        haveSwaps = False
        while right_point < endpoint:
            if array[left_point] > array[right_point]:
                array[left_point], array[right_point] = array[right_point], array[left_point]
                haveSwaps = True
            right_point += 1
            left_point += 1
        endpoint -= 1  
    return

def main():
    print("Here is the shuffled list: ", my_list)
    bubble_sort(my_list)
    print("Here is the bubble sorted list: ", my_list)
    return 

if __name__ == "__main__":
    main()

