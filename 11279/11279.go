package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Heap []int

func (h Heap)Len() int {return len(h)}
func (h Heap)Less(i, j int) bool {return h[i] > h[j]}
func (h Heap)Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *Heap)Push(num interface{}) {
	*h = append(*h, num.(int))
}
func (h *Heap)Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)

	MaxHeap := &Heap{}
	heap.Init(MaxHeap)
	for i:=0 ; i<N ; i++ {
		var input int
		fmt.Fscan(reader, &input)
		if input == 0 {
			if len(*MaxHeap) == 0 {
				fmt.Fprintln(writer, 0)
			} else {
				fmt.Fprintln(writer, heap.Pop(MaxHeap))
			}
		}else {
			heap.Push(MaxHeap, input)
		}
	}
}
