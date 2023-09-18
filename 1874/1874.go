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

	arr := make([]int,N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	stack := list.New()
	index := 0
	answer := []string{}
	for i:=1 ; i<=N ; i++ {
		if i <= arr[index] {
			stack.PushBack(i)
			answer = append(answer, "+")
		}
		for index < N && stack.Len() > 0{
			if stack.Back().Value == arr[index] {
				answer = append(answer, "-")
				stack.Remove(stack.Back())
				index++
			} else {
				break
			}
		}
	}
	if stack.Len() != 0 {
		fmt.Fprintln(writer, "NO")
		return
	}
	for i:=0 ; i<len(answer) ; i++ {
		fmt.Fprintln(writer, answer[i])
	}
}