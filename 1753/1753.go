package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type link struct {
	to int
	value int
}

type minHeap []link

func (h minHeap) Len() int {return len(h)}
func (h minHeap) Less(i, j int) bool {return h[i].value < h[j].value}
func (h minHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *minHeap) Push(input interface{}) {
	*h = append(*h, input.(link))
}
func (h *minHeap) Pop() interface{} {
	length := len(*h)
	old := *h
	value := old[length-1]
	*h = old[:length-1]
	return value
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	input := strings.Split(Input(scanner), " ")
	N, _ := strconv.Atoi(input[0])
	M, _ := strconv.Atoi(input[1])

	start, _ := strconv.Atoi(Input(scanner))

	linkList := make(map[int][]link)
	for i:=0 ; i<M ; i++ {
		linkInput := strings.Split(Input(scanner), " ")
		from, _ := strconv.Atoi(linkInput[0])
		to, _ := strconv.Atoi(linkInput[1])
		value, _ := strconv.Atoi(linkInput[2])
		linkList[from] = append(linkList[from], link{to, value})
	}

	result := Process(N, start, linkList)
	for _, v := range result {
		fmt.Fprintln(writer, v)
	}
}

func Process(N, start int, linkList map[int][]link) []string {
	checkValue := make(map[int]int)
	for i:=1 ; i<=N ; i++ {
		checkValue[i] = 1<<30
	}
	list := &minHeap{}
	checkValue[start] = 0
	heap.Init(list)
	heap.Push(list, link{start, 0})
	for list.Len() > 0 {
		now := heap.Pop(list).(link)
		if now.to != start && checkValue[now.to] < now.value {
			continue
		}
		for i:=0 ; i<len(linkList[now.to]) ; i++ {
			next := linkList[now.to][i]
			if checkValue[next.to] > checkValue[now.to]+next.value {
				checkValue[next.to] = checkValue[now.to]+next.value
				heap.Push(list, link{next.to, checkValue[next.to]})
			}
		}
	}
	result := []string{}
	for i:=1 ; i<=N ; i++ {
		if checkValue[i] == 1<<30 {
			result = append(result, "INF")
		}else {
			result = append(result, strconv.Itoa(checkValue[i]))
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Input(scanner *bufio.Scanner) string {
	scanner.Scan()

	return scanner.Text()
}