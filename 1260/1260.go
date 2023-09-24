package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)
func main(){

	defer writer.Flush()

	var N, M, V int
	fmt.Fscan(reader, &N, &M, &V)

	link := make(map[int][]int)
	checkDFS := make(map[int]bool)
	checkBFS := make(map[int]bool)

	for i:=0 ; i<M ; i++ {
		var A, B int
		fmt.Fscan(reader, &A, &B)
		link[A] = append(link[A], B)
		link[B] = append(link[B], A)
		
	}

	queue := []int{V}
	stack := []int{V}

	for i:=1 ; i<=N ; i++ {
		sort.Ints(link[i])
	}

	DFS(link, queue, checkDFS)
	fmt.Fprint(writer, "\n")
	BFS(link, stack, checkBFS)
	
}
func DFS(link map[int][]int, stack []int, check map[int]bool){
	
	for len(stack) > 0 {
		length := len(stack)
		nowNode := stack[length-1]
		stack = stack[:(length-1)]

		if check[nowNode] == true {
			continue
		}

		check[nowNode] = true
		fmt.Fprintf(writer, "%d ",nowNode)

		for i:=1 ; i<=len(link[nowNode]) ; i++ {
			v := link[nowNode][len(link[nowNode])-i]
			if check[v] == false {
				stack = append(stack, v)
			}
		}
	}
}

func BFS(link map[int][]int, queue []int, check map[int]bool){
	for len(queue) > 0 {
		nowNode := queue[0]
		queue = queue[1:]

		if check[nowNode] == true {
			continue
		}

		check[nowNode] = true
		fmt.Fprintf(writer, "%d ",nowNode)

		for _, v := range link[nowNode] {
			if check[v] == false {
				queue = append(queue, v)
			}
		}
	}
}