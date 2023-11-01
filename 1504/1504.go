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
	length int
}

type minHeap []link

func (h minHeap) Len() int {return len(h)}
func (h minHeap) Less(i, j int) bool {return h[i].length<h[j].length}
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
	linkList := make(map[int][]link)
	for i:=0 ; i<M ; i++ {
		linkInput := strings.Split(Input(scanner), " ")
		from, _ := strconv.Atoi(linkInput[0])
		to, _ := strconv.Atoi(linkInput[1])
		length, _ := strconv.Atoi(linkInput[2])
		linkList[from] = append(linkList[from], link{to, length})
		linkList[to] = append(linkList[to], link{from, length})
	}
	nodeList := strings.Split(Input(scanner), " ")
	node1, _ := strconv.Atoi(nodeList[0])
	node2, _ := strconv.Atoi(nodeList[1])
	result1 := findRoute(1, node1, node2, N, N, linkList)
	result2 := findRoute(1, node2, node1, N, N, linkList)
	fmt.Fprintln(writer, min(result1, result2))
}

func findRoute(from, node1, node2, to, N int, linkList map[int][]link) int {
	first := calcLength(from, node1, N, linkList)
	middle := calcLength(node1, node2, N, linkList)
	last := calcLength(node2, to, N, linkList)

	if first == -1 || middle == -1 || last == -1 {
		return -1
	}
	return first+middle+last
}

func calcLength(from, to, N int, linkList map[int][]link) int {
	if from == to {
		return 0
	}
	list := &minHeap{}
	checkLength := make(map[int]int)
	for i:=1 ; i<=N ; i++ {
		checkLength[i] = 800*1000*3
	}
	checkLength[from] = 0
	isVisit := make(map[int]bool)
	isVisit[from] = true
	heap.Init(list)
	for _, v := range linkList[from] {
		heap.Push(list, v)
		checkLength[v.to] = min(checkLength[v.to], v.length)
		isVisit[v.to] = true
	}
	
	for list.Len() > 0{
		now := heap.Pop(list).(link)
		if checkLength[now.to] < now.length {
			continue
		}
		for i:=0 ; i<len(linkList[now.to]) ; i++ {
			next := linkList[now.to][i]
			if checkLength[next.to] > next.length+checkLength[now.to] {
				checkLength[next.to] = next.length+checkLength[now.to]
				heap.Push(list, next)
				isVisit[next.to] = true
			}
		}
	}
	if isVisit[to] == false {
		return -1
	}
	return checkLength[to]
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