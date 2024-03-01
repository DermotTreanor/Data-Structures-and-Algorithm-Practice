package main

import (
	"fmt"
)

func anagrams(word string)(anagram []string) {
	//Base case is if we have the 'word' containing just one lette.
	//The collection of anagrams for this is just the single entry of the letter itself.
	if len(word) == 1{
		return []string{word}
	}

	//Make an empty list for the current collection of anagrams.
	listy := []string{}

	//Cycle through all the anagrams we have for everything but the last letter (and in the process we use recursion to make them)
	for _, list_item := range anagrams(word[:len(word) - 1]){

		//For each anagram we add the final letter to every possible position. That makes a new anagram we add to our list.
		for string_index := 0; string_index <= len(list_item); string_index++ { //Notice we go to the len value (which is one index more than our curent set of anagrams). 
																				//That will be so we run this up to and including the last space so we can add the letter to the end too.
			new_string := list_item[:string_index] + string(word[len(word) - 1]) + list_item[string_index:]
			listy = append(listy, new_string)
		}
	}

	return listy
}

func main() {
	for _, v := range anagrams("abcd"){
		fmt.Println(v)
	}
}
