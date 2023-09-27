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

	friendGraph := make([][]int,N+1)
	for i:=0 ; i<M ; i++ {
		var A, B int
		fmt.Fscan(reader, &A, &B)
		friendGraph[A] = append(friendGraph[A], B)
		friendGraph[B] = append(friendGraph[B], A)
	}
	cnt := 1
	min := N*N
	answer := 0
	for i:=1 ; i<=N ; i++ {
		check := make(map[int]bool)
		check[i] = true
		queue := []int{i}
		result := BFS(friendGraph, queue, cnt, check)
		if min > result {
			min = result
			answer = i
		}
	}

	fmt.Fprintln(writer, answer)
}
func BFS(friendGraph [][]int, queue []int, cnt int, check map[int]bool) int {
	if len(queue) == 0 {
		return 0
	}
	newQueue := []int{}
	answer := 0 
	for len(queue) > 0 {
		now := queue[0]
		queue = queue[1:]
		for i:=0 ; i<len(friendGraph[now]) ; i++ {
			if check[friendGraph[now][i]] == false {
				newQueue = append(newQueue, friendGraph[now][i])
				check[friendGraph[now][i]] = true
				answer += cnt
			}
		}
	}
	answer += BFS(friendGraph, newQueue, cnt+1, check)
	return answer
}