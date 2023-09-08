package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)
	queue := []int{}
	for i:=0 ; i<N ; i++{
		var command string
		fmt.Fscan(reader, &command)
		switch command {
		case "push" :
			var Num int
			fmt.Fscan(reader, &Num)
			queue = append(queue, Num)
		case "pop" :
			if len(queue) > 0 {
				fmt.Fprintln(writer, queue[0])
				queue = queue[1:]
			}else {
				fmt.Fprintln(writer, -1)
			}
		case "size" :
			fmt.Fprintln(writer, len(queue))
		case "empty" :
			if len(queue) > 0 {
				fmt.Fprintln(writer, 0)
			}else {
				fmt.Fprintln(writer, 1)
			}
		case "front" :
			if len(queue) > 0 {
				fmt.Fprintln(writer, queue[0])
			}else {
				fmt.Fprintln(writer, -1)
			}
		case "back" :
			if len(queue) > 0 {
				fmt.Fprintln(writer, queue[len(queue)-1])
			}else {
				fmt.Fprintln(writer, -1)
			}
		}
	}
}