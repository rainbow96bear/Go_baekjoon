package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)
type MinHeap []int
type MaxHeap []int

func (h MinHeap)Len() int {return len(h)}
func (h MinHeap)Less(i, j int) bool {return h[i]<h[j]}
func (h MinHeap)Swap(i, j int) { h[i], h[j] = h[j], h[i]}
func (h *MinHeap)Push(num interface{}) {
	*h = append(*h, num.(int))
}
func (h *MinHeap)Pop() interface{} {
	old := *h
	length := len(*h)
	num := old[length-1]
	*h = old[:length-1]
	return num
}

func (h MaxHeap)Len() int {return len(h)}
func (h MaxHeap)Less(i, j int) bool {return h[i]>h[j]}
func (h MaxHeap)Swap(i, j int) { h[i], h[j] = h[j], h[i]}
func (h *MaxHeap)Push(num interface{}) {
	*h = append(*h, num.(int))
}
func (h *MaxHeap)Pop() interface{} {
	old := *h
	length := len(*h)
	num := old[length-1]
	*h = old[:length-1]
	return num
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)

	leftArr := &MaxHeap{}
	rightArr := &MinHeap{}
	heap.Init(leftArr)
	heap.Init(rightArr)
	middleValue := 0
	for i:=0 ; i<N ; i++ {
		var input int
		fmt.Fscan(reader, &input)
		if i == 0 {
			middleValue = input
		}else {
			if middleValue <= input {
				heap.Push(rightArr, input)
			}else {
				heap.Push(leftArr, input)
			}
			if leftArr.Len() + 2 <= rightArr.Len() {
				heap.Push(leftArr, middleValue)
				middleValue = (heap.Pop(rightArr)).(int)
			}else if leftArr.Len() >= rightArr.Len() + 1{
				heap.Push(rightArr, middleValue)
				middleValue = (heap.Pop(leftArr)).(int)
			}
		}
		fmt.Fprintln(writer, middleValue)
	}
}