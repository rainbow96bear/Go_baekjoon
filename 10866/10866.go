package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)
func main(){

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)
	deque := list.New()
	for i:=0 ; i<N ; i++ {
		var command string
		var input int
		fmt.Fscan(reader, &command)
		isPush := command == "push_front" || command == "push_back"
		if isPush {
			fmt.Fscan(reader, &input)
		}
		process(deque, command, input)
	}
}

func process(deque *list.List, command string, input int){
	length := deque.Len()

	switch command {
	case "push_front" :
		deque.PushFront(input)
	case "push_back" :
		deque.PushBack(input)
	case "pop_front" :
		if length == 0 {
			fmt.Fprintln(writer, -1)
			return
		}
		fmt.Fprintln(writer, deque.Front().Value)
		deque.Remove(deque.Front())
	case "pop_back" :
		if length == 0 {
			fmt.Fprintln(writer, -1)
			return
		}
		fmt.Fprintln(writer, deque.Back().Value)
		deque.Remove(deque.Back())
	case "size" :
		fmt.Fprintln(writer, length)
	case "empty" :
		if length == 0 {
			fmt.Fprintln(writer, 1)
			return
		}
		fmt.Fprintln(writer, 0)
	case "front" :
		if length == 0 {
			fmt.Fprintln(writer, -1)
			return
		}
		fmt.Fprintln(writer, deque.Front().Value)
	case "back" :
		if length == 0 {
			fmt.Fprintln(writer, -1)
			return
		}
		fmt.Fprintln(writer, deque.Back().Value)
	}
}