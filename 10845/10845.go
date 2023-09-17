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

	queue := list.New()
	for i:=0 ; i<N ; i++ {
		var command string
		var Num int
		fmt.Fscan(reader, &command)
		if command == "push" {
			fmt.Fscan(reader, &Num)
		}
		process(queue, command, Num)
	}
}

func process(queue *list.List, command string, Num int){
	length := queue.Len()

	switch command {
	case "push" : 
		queue.PushBack(Num)
	case "pop" :
		if length == 0 {
			fmt.Fprintln(writer, -1)
			return
		}
		fmt.Fprintln(writer, queue.Front().Value)
		queue.Remove(queue.Front())
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
		fmt.Fprintln(writer, queue.Front().Value)
	case "back" :
		if length == 0 {
			fmt.Fprintln(writer, -1)
			return
		}
		fmt.Fprintln(writer, queue.Back().Value)
	}
}