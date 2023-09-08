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
	queue := make([]int,N)
	stack := []int{}
	answer := "Nice"
	for i:=0 ; i<N ; i++{
		fmt.Fscan(reader, &queue[i])
	}
	cnt := 1
	for len(queue)>0{
		if cnt != queue[0] {
			if len(stack) > 0 && stack[len(stack)-1] == cnt {
				stack = stack[:(len(stack)-1)]
				cnt++
			} else {
				stack = append(stack, queue[0])
				queue = queue[1:]
			}
		} else {
			queue = queue[1:]
			cnt++
		}
	}
	for len(stack) > 0{
		if cnt == stack[len(stack)-1]{
			stack = stack[:(len(stack)-1)]
			cnt++
		} else {
			answer = "Sad"
			break
		}
	}
	fmt.Fprint(writer, answer)
}