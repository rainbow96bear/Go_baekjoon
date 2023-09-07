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

	var K int
	fmt.Fscan(reader, &K)

	stack := []int{}
	for i:=0 ; i<K ; i++{
		var input int
		fmt.Fscan(reader, &input)
		if input == 0 {
			stack = stack[:(len(stack)-1)]
		}else{
			stack = append(stack, input)
		}
	}
	answer := 0
	for _, v := range stack {
		answer += v
	}
	fmt.Fprint(writer, answer)
}