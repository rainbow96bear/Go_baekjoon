package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)
type Quiz struct {
	number int
	value int
}

type List []Quiz

func (h List) Len() int {return len(h)}
func (h List) Less(i, j int) bool {return h[i].number < h[j].number}
func (h List) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *List) Push(x interface{}) {
	*h = append(*h, x.(Quiz))
}
func (h *List) Pop() interface{} {
	old := *h
	length := len(old)
	result := old[length-1]
	*h = old[:length-1]
	return result
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N, M int
	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	QuizList := make([]Quiz, N+1)
	Relationship := make([][]int, N+1)
	for i:=1 ; i<=N ; i++ {
		QuizList[i].number = i
	}
	for i:=1 ; i<=M ; i++ {
		var first, last int
		fmt.Fscanf(reader, "%d %d\n", &first, &last)
		Relationship[first] = append(Relationship[first], last)
		QuizList[last].value++
	}
	result := FindOrder(N, M, QuizList, Relationship)
	for _, v := range result {
		fmt.Fprintf(writer, "%d ", v)
	}
}

func FindOrder(N, M int, QuizList []Quiz, Relationship [][]int) []int {
	list := &List{}
	heap.Init(list)

	for i:=1 ; i<=N ; i++ {
		if QuizList[i].value == 0 {
			heap.Push(list, QuizList[i])
		}
	}
	result := []int{}
	for list.Len() > 0 {
		now := heap.Pop(list).(Quiz)
		result = append(result, now.number)
		for _, v := range Relationship[now.number] {
			QuizList[v].value--
			if QuizList[v].value == 0 {
				heap.Push(list,  QuizList[v])
			}
		}
	}
	return result
}