package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type PossibleList []jewelry

type jewelry struct {
	M int
	V int
}

func (p PossibleList) Len() int           { return len(p) }
func (p PossibleList) Less(i, j int) bool { return p[i].V > p[j].V }
func (p PossibleList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *PossibleList) Push(x interface{}) {
	item := x.(jewelry)
	*p = append(*p, item)
}
func (p *PossibleList) Pop() interface{} {
	old := *p
	length := len(old)
	result := old[length-1]
	*p = old[:length-1]
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N, K int
	fmt.Fscanf(reader, "%d %d\n", &N, &K)

	JewelryList := make([]jewelry, N)
	for i := 0; i < N; i++ {
		fmt.Fscanf(reader, "%d %d\n", &JewelryList[i].M, &JewelryList[i].V)
	}

	BagList := make([]int, K)
	for i := 0; i < K; i++ {
		fmt.Fscanf(reader, "%d\n", &BagList[i])
	}

	result := FindMaxValue(N, K, JewelryList, BagList)
	fmt.Fprintln(writer, result)
}

func FindMaxValue(N, K int, JewelryList []jewelry, BagList []int) int64 {
	sort.Slice(JewelryList, func(i, j int) bool { return JewelryList[i].M < JewelryList[j].M })
	sort.Slice(BagList, func(i, j int) bool { return BagList[i] < BagList[j] })

	possibleList := &PossibleList{}
	heap.Init(possibleList)
	totalValue := int64(0)

	idx := 0
	for i := 0; i < K; i++ {
		BagMax := BagList[i]
		for idx < N && JewelryList[idx].M <= BagMax {
			heap.Push(possibleList, JewelryList[idx])
			idx++
		}
		if possibleList.Len() > 0 {
			jewelry := heap.Pop(possibleList).(jewelry)
			totalValue += int64(jewelry.V)
		}
	}
	return totalValue
}
