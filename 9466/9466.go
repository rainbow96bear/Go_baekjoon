package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAX = 100001

var (
	cnt 		int
	graph     [MAX]int
	visited   [MAX]bool
	done      [MAX]bool
)

func hasCycle(node int) {
	visited[node] = true
	next := graph[node]

	if !visited[next] {
		hasCycle(next)
	} else if !done[next] {
		for i := next; i != node; i = graph[i] {
			cnt++
		}
		cnt++
	}
	done[node] = true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()
	var t int
	fmt.Fscanf(reader, "%d\n", &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscanf(reader, "%d\n", &n)

		cnt = 0

		for i := 1; i <= n; i++ {
			fmt.Fscanf(reader, "%d", &graph[i])
		}
		fmt.Fscanf(reader, "\n")
		for i := 1; i <= n; i++ {
			if !visited[i] {
				hasCycle(i)
			}
		}

		fmt.Fprintln(writer, n - cnt)
		cnt = 0
		resetArrays()
	}
}

func resetArrays() {
	for i := range visited {
		visited[i] = false
	}
	for i := range done {
		done[i] = false
	}
}