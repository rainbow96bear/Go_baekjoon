package main

import (
	"bufio"
	"fmt"
	"os"
)
var (
	tree = make([]int, 100001)
)
func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N, M int
	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	for i:=1 ; i<=N ; i++ {
		tree[i] = i
	}
	for i:=0 ; i<M ; i++ {
		var building1, building2 int
		fmt.Fscanf(reader, "%d %d\n", &building1, &building2)
		Union(building1, building2)
	}

	schedule := make([]int, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscanf(reader, "%d ", &schedule[i])
	}

	answer := 0
	now := schedule[0]
	for i:=1 ; i<N ; i++ {
		if Find(tree[now]) != Find(tree[schedule[i]]) {
			answer++
		}
		now = schedule[i]
	}
	fmt.Fprintln(writer, answer)
}

func Find(building int) int {
	if tree[building] == building {
		return building
	}
	tree[building] = Find(tree[building])
	return tree[building]
}

func Union(building1, building2 int) {
	if Find(building1) > Find(building2) {
		tree[Find(building1)] = Find(building2)
	}else {
		tree[Find(building2)] = Find(building1)
	}
}