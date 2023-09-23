package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    var N int
    fmt.Scan(&N)

    minHeap := &MinHeap{}
    heap.Init(minHeap)

    for i := 0; i < N; i++ {
        var X int
        scanner.Scan()
        X, _ = strconv.Atoi(scanner.Text())

        if X == 0 {
            if minHeap.Len() == 0 {
                fmt.Fprintln(writer, 0)
            } else {
                min := heap.Pop(minHeap).(int)
                fmt.Fprintln(writer, min)
            }
        } else {
            heap.Push(minHeap, X)
        }
    }
}
