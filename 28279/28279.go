package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)

	deque := list.New()
	for i:=0 ; i<N ; i++{
		var command int
		fmt.Fscan(reader, &command)

		switch command {
		case 1 :
			var input int
			fmt.Fscan(reader, &input)
			deque.PushFront(input)
		case 2 :
			var input int
			fmt.Fscan(reader, &input)
			deque.PushBack(input)
		case 3 :
			if deque.Len() > 0 {
				fmt.Fprintln(writer, deque.Front().Value)
				deque.Remove(deque.Front())
			}else {
				fmt.Fprintln(writer, -1)
			}
		case 4 :
			if deque.Len() > 0 {
				fmt.Fprintln(writer, deque.Back().Value)
				deque.Remove(deque.Back())
			}else {
				fmt.Fprintln(writer, -1)
			}
		case 5 :
			fmt.Fprintln(writer, deque.Len())
		case 6 :
			if deque.Len() > 0 {
				fmt.Fprintln(writer, 0)
			}else {
				fmt.Fprintln(writer, 1)
			}
		case 7 :
			if deque.Len() > 0 {
				fmt.Fprintln(writer, deque.Front().Value)
			}else {
				fmt.Fprintln(writer, -1)
			}
		case 8 :
			if deque.Len() > 0 {
				fmt.Fprintln(writer, deque.Back().Value)
			}else {
				fmt.Fprintln(writer, -1)
			}
		}
	}
}