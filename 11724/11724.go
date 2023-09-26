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

	var N, M int
	fmt.Fscan(reader, &N, &M)

	graph := make([][]int, N+1)
	for i:=0 ; i<M ; i++ {
		var A, B int
		fmt.Fscan(reader, &A, &B)
		graph[A]= append(graph[A], B)
		graph[B]= append(graph[B], A)
	}

	check := make(map[int]bool)
	cnt := 0
	for i:=1 ; i<=N ; i++ {
		if check[i] == false {
			cnt++
			stack := []int{i}
			for len(stack)>0 {
				nowNode := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				for _, v := range graph[nowNode] {
					if check[v] == false {
						stack = append(stack, v)
						check[v]=true
					}
				}
			}
		}
	}
	fmt.Fprintln(writer, cnt)
}