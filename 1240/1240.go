package main

import (
	"bufio"
	"fmt"
	"os"
)

type Child struct {
	node int
	length int
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N, M int
	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	tree := make([][]Child, N+1)
	for i:=1 ; i<N ; i++ {
		var a, b, length int
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &length)
		tree[a] = append(tree[a], Child{b, length})
		tree[b] = append(tree[b], Child{a, length})
	}

	isVisit := make([]bool, N+1)
	for i:=0 ; i<M ; i++ {
		var from, to int
		fmt.Fscanf(reader, "%d %d\n", &from, &to)
		result := dfs(from, to, tree, isVisit)
		fmt.Fprintln(writer, result)
	}
}

func dfs(from, to int, tree [][]Child, isVisit []bool) int {
	isVisit[from] = true
	length := 0
	for _, v := range tree[from] {
		if v.node == to {
			length = v.length
			break
		}
		if isVisit[v.node] == false {
			result := dfs(v.node, to, tree, isVisit)
			if result != 0 {
				length = v.length + result
				break
			}
		}
	}
	isVisit[from] = false
	return length
}