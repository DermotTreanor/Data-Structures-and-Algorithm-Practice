/*
This is a basic implementation and demonstration of a stack in GO. 
It will accept any value into its collection of 'entities' and we can
push or pop onto or off the stack.
*/
package main
import "fmt"

type stack struct {
	entities []interface{}
}
func (s *stack) push(i interface{}) {
	s.entities = append(s.entities, i)
}
func (s *stack) pop() (v interface{}) {
	if len(s.entities) > 0 {
		v = s.entities[len(s.entities)-1]
		s.entities = s.entities[:len(s.entities)-1]
	}
	return v
}

func main() {
	var myStack stack
	//Adding different data types
	myStack.push(false)
	myStack.push("hello")
	myStack.push(0)
	myStack.push("goodbye")
	myStack.push(true)
	myStack.push(90.84)
	myStack.push(44032314)
	myStack.push(nil)
	myStack.push('\x07')
	fmt.Println("\u001b[38;5;43mInitial stack entities:", "\u001b[38;5;84m", myStack.entities, "\u001b[0m")

	//Popping the last three
	fmt.Println("\tPopping:", myStack.pop())
	fmt.Println("\tPopping:", myStack.pop())
	fmt.Println("\tPopping:", myStack.pop())
	fmt.Println("\tPopping:", myStack.pop())

	fmt.Println("\u001b[38;5;45mRemaining stack entities:", "\u001b[38;5;45m", myStack.entities, "\u001b[0m")
}
