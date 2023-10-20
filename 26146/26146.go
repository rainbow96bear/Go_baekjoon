package main

import (
	"bufio"
	"fmt"
	"os"
)

var visited []bool
var stack []int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	
	defer writer.Flush()

	var N, M int
	fmt.Fscan(reader, &N, &M)

	visited = make([]bool, N+1)
	graph := make([][]int, N+1)
	reversegraph := make([][]int, N+1)

	for i := 1; i <= N; i++ {
		graph[i] = make([]int, 0)
		reversegraph[i] = make([]int, 0)
	}

	for i:=0; i<M; i++ {
		var from, to int
		fmt.Fscan(reader, &from, &to)

		graph[from] = append(graph[from], to)
		reversegraph[to] = append(reversegraph[to], from)
	}

	for i := 1; i <= N; i++ {
		if visited[i] == false {
			dfs(i, graph, true)
		}
	}

	visited = make([]bool, N+1)

	now := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	dfs(now, reversegraph, false)

	answer := "Yes"
	for len(stack) > 0 {
		if visited[stack[len(stack)-1]] == false {
			answer = "No"
			break
		}
		stack = stack[:len(stack)-1]
	}

	fmt.Print(answer)
}

func dfs(depth int, graph [][]int, check bool) {
	visited[depth] = true

	for _, a := range graph[depth] {
		if !visited[a] {
			dfs(a, graph, check)
		}
	}

	if check {
		stack = append(stack, depth)
	}
}
