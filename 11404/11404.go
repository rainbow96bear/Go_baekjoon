package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bus struct {
	to int
	price int
}

type BusList []Bus

func (h BusList)Len() int {return len(h)} 
func (h BusList)Less(i, j int) bool {return h[i].price<h[j].price}
func (h BusList)Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *BusList)Push(bus interface{}) {
	*h = append(*h, bus.(Bus))
}
func (h *BusList)Pop() interface{} {
	length := len(*h)
	old := *h
	bus := old[length-1]
	*h = old[:length-1]
	return bus
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	n, _ := strconv.Atoi(Input(scanner))

	m, _ := strconv.Atoi(Input(scanner))

	busList := make(map[int][]Bus)
	for i:=0 ; i<m ; i++ {
		input := strings.Split(Input(scanner), " ")
		start, _ := strconv.Atoi(input[0])
		end, _ := strconv.Atoi(input[1])
		price, _ := strconv.Atoi(input[2])
		busList[start] = append(busList[start], Bus{end, price})
	}
	answer := Process(n, busList)
	for i:=0 ; i<n ; i++ {
		for j:=0 ; j<n ; j++ {
			fmt.Fprintf(writer, "%d ", answer[i][j])
		}
		fmt.Fprintln(writer)
	}
}

func Process(n int, busList map[int][]Bus) [][]int {
	answer := make([][]int, n)
	for i:=1 ; i<=n ; i++ {
		answer[i-1] = findRoute(i, n, busList)
	}
	return answer
}

func findRoute(start, n int, busList map[int][]Bus) []int {
	checkPrice := make(map[int]int)
	for i:=1 ; i<=n ; i++ {
		checkPrice[i] = 1<<30
	}
	list := &BusList{}
	answer := make([]int, n)
	heap.Init(list)
	checkPrice[start] = 0
	heap.Push(list, Bus{start, 0})
	for list.Len() > 0 {
		bus := heap.Pop(list).(Bus)
		now := bus.to
		for _, v := range busList[now] {
			if checkPrice[v.to] > checkPrice[now] + v.price {
				answer[v.to-1] = checkPrice[now] + v.price
				checkPrice[v.to] = checkPrice[now] + v.price
				heap.Push(list, v)
			}
		}
	}
	return answer
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Input(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}