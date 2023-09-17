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

	stack := list.New()
	for i:=0 ; i<N ; i++ {
		var command string
		var Num int
		fmt.Fscan(reader, &command)
		if command == "push"{
			fmt.Fscan(reader, &Num)
		}
		process(stack, command, Num)
	}
}

func process(stack *list.List, command string, Num int){
	length := stack.Len()
	switch command {
	case "push" :
		stack.PushBack(Num)
	case "pop" :
		if length != 0 {
			fmt.Fprintln(writer, stack.Back().Value)
			stack.Remove(stack.Back())
			return
		}
		fmt.Fprintln(writer, -1)
	case "size" :
		fmt.Fprintln(writer, length)
	case "empty" :
		if length != 0 {
			fmt.Fprintln(writer, 0)
			return
		}
		fmt.Fprintln(writer, 1)
	case "top" :
		if length != 0 {
			fmt.Fprintln(writer, stack.Back().Value)
			return
		}
		fmt.Fprintln(writer, -1)
	}
}