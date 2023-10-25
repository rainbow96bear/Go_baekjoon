package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)
type bus struct {
	to int
	cost int
}

type MinHeap []bus

func (h MinHeap) Len() int {return len(h)}
func (h MinHeap) Less(i, j int) bool {return h[i].cost < h[j].cost}
func (h MinHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *MinHeap) Push(value interface{}) {
	*h = append(*h, value.(bus))
}
func (h *MinHeap) Pop() interface{} {
	length := len(*h)
	old := *h
	value := old[length-1]
	*h = old[:length-1]
	return value
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)

	var M int
	fmt.Fscan(reader, &M)

	link := make(map[int][]bus)

	for i:=0 ; i<M ; i++ {
		var from, to, cost int
		fmt.Fscan(reader, &from, &to, &cost)
		link[from] = append(link[from], bus{to, cost})
	}
	var start, end int
	fmt.Fscan(reader, &start, &end)
	answer := Process(link, start, end, N)
	fmt.Fprintln(writer, answer)
}

func Process(link map[int][]bus, start, end, N int) int {
    queue := &MinHeap{}

	heap.Init(queue)
	heap.Push(queue, bus{start, 0})

	distance := make(map[int]int)

	for i:=1 ; i<=N ; i++ {
		distance[i] = 1000000000
	}

	distance[start] = 0
	for queue.Len() > 0 {
		now := heap.Pop(queue).(bus)

		if now.cost > distance[now.to] {
			continue
		}
		for i := 0; i < len(link[now.to]); i++ {
			nextNode := link[now.to][i]
			if distance[now.to] + nextNode.cost < distance[nextNode.to] {
				distance[nextNode.to] = min(distance[now.to] + nextNode.cost, distance[nextNode.to])
				heap.Push(queue, nextNode)
			}
		}
	}
	return distance[end]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}