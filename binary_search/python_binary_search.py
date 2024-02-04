#! /usr/bin/python3
import os
import sys


def binary_search(array, x):
    start = 0 
    end = len(array) - 1
    
    while start <= end: #Make sure we are running the loop also when they are equal. 
                        #We only want to end the loop when they go passed each other, i.e. when start is more than end
        midpoint = start + ((end - start) // 2)
        if x == array[midpoint]:
            return True, midpoint
        elif x < array[midpoint]:
            end = midpoint - 1
        elif x > array[midpoint]:
            start = midpoint + 1 

    return False, None


def main():
    COLOURS = {
    'GREEN': '\033[32m',
    'RED': '\033[31m',
    'WHITE': '\033[37m',
    'YELLOW': '\033[33m'
    }
    os.system('clear')
    my_list = [2, 4, 6, 8, 10, 12, 24, 36, 64, 128, 256, 512, 1024, 2048]
    print(f"Running {COLOURS['YELLOW']}{sys.argv[0]}{COLOURS['WHITE']} file...")
    print("This is the ordered list", my_list)


    value_to_find = int(input("Please enter a value to search for: "))
    presence, index = binary_search(my_list, value_to_find)
    if presence:
        print(f"The value {COLOURS['GREEN']}IS{COLOURS['WHITE']} in the list. It is at the following index: ", index)
    else:
        print(f"The value {COLOURS['RED']}IS NOT{COLOURS['WHITE']} in the list.")


print(__name__)
if __name__ == "__main__":
    main()
