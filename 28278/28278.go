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
	
	stack := []int{}
	for i:=0 ; i<N ; i++{
		var command int
		fmt.Fscan(reader, &command)

		switch command {
			case 1 :
				var input int
				fmt.Fscan(reader, &input)
				stack = append(stack, input)
			case 2 :
				if len(stack) > 0	{
					fmt.Fprintln(writer, stack[len(stack)-1])
					stack = stack[:(len(stack)-1)]
				}else {
					fmt.Fprintln(writer, -1)
				}
			case 3 :
				fmt.Fprintln(writer, len(stack))
			case 4 :
				if len(stack) > 0 {
					fmt.Fprintln(writer, 0)
				}else {
					fmt.Fprintln(writer, 1)
				}
			case 5 :
				if len(stack) > 0 {
					fmt.Fprintln(writer, stack[len(stack)-1])
				}else {
					fmt.Fprintln(writer, -1)
				}
		}
	}
}