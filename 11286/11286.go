package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return abs(h[i]) < abs(h[j]) } // 절댓값으로 비교
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    ph := &IntHeap{}
    mh := &IntHeap{}

    heap.Init(ph)
    heap.Init(mh)

    var N int
    fmt.Fscan(reader, &N)

    for i := 0; i < N; i++ {
        var input int
        fmt.Fscan(reader, &input)

        if input != 0 {
            if input > 0 {
                heap.Push(ph, input)
            } else {
                heap.Push(mh, input)
            }
        } else {
            if ph.Len() == 0 && mh.Len() == 0 {
                fmt.Fprintln(writer, 0)
            } else if ph.Len() == 0 {
                fmt.Fprintln(writer, heap.Pop(mh))
            } else if mh.Len() == 0 {
                fmt.Fprintln(writer, heap.Pop(ph))
            } else {
                plusMin := heap.Pop(ph).(int)
                minusMin := heap.Pop(mh).(int)
                if abs(plusMin) < abs(minusMin) {
                    fmt.Fprintln(writer, plusMin)
                    heap.Push(mh, minusMin)
                } else {
                    fmt.Fprintln(writer, minusMin)
                    heap.Push(ph, plusMin)
                }
            }
        }
    }
}