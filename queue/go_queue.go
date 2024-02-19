package main

import "fmt"

type Queue struct{
	collection []interface{}
}

//Using variadic parameters for educational purposes 
func (q *Queue)enqueue(i ...interface{}){
	q.collection = append(q.collection, i...)
}

func (q *Queue)dequeue()(value interface{}){
	if len(q.collection) > 0{
		value = q.collection[0]
		q.collection = q.collection[1:]
		return value
	} else{
		//May need to consider returning both a value and an error.
		//When the queue is empty it will look like it's just full of nil 'values' (since
		//we allow any type in the collection this isn't illogical)
		return nil
	}
}

func main(){
	myQueue := Queue{}

	
	//Enqueue a number of items.
	myQueue.enqueue("hello", "how", "are", 37, true, "YOU", nil, 'g', false)
	fmt.Println(myQueue.collection)
	myQueue.enqueue(19001237)
	fmt.Println(myQueue.collection)
	fmt.Println("\n")

	//Get items out of the front of the Queue
	fmt.Println("Dequeue: ", myQueue.dequeue())
	fmt.Println(myQueue.collection)

	fmt.Println("Dequeue: ", myQueue.dequeue())
	fmt.Println(myQueue.collection)

	fmt.Println("Dequeue: ", myQueue.dequeue())
	fmt.Println(myQueue.collection)

}