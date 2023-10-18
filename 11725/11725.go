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
	link := make(map[int][]int)
	for i:=0 ; i<N-1 ; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		link[a] = append(link[a], b)
		link[b] = append(link[b], a)
	}

	Process(link, N, writer)
}

func Process(link map[int][]int, N int, writer *bufio.Writer){

	stack := []int{1}
	checkParent := make(map[int]int)
	checkParent[1] = 1
	for len(stack) > 0 {
		length := len(stack)
		now := stack[length-1]
		stack = stack[:length-1]
		for i:=0 ; i<len(link[now]) ; i++ {
			if checkParent[link[now][i]] == 0 {
				checkParent[link[now][i]] = now
				stack = append(stack, link[now][i])
			}
		}
	}
	for i:=2 ; i<=N ; i++ {
		fmt.Fprintln(writer, checkParent[i])
	}
}