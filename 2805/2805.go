package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N, M int
	fmt.Fscan(reader, &N, &M)

	treeList := make([]int, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscan(reader, &treeList[i])
	}

	sort.Ints(treeList)
	min := 0
	max := treeList[len(treeList)-1]
	mid := 0
	for min <= max {
		mid = (min+max) / 2
		result := cutting(treeList, mid)
		
		if result >= M {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	fmt.Fprintln(writer, max)
}

func cutting(treeList []int, height int) int{
	sum := 0
	for i:=0 ; i<len(treeList) ; i++ {
		if treeList[i] > height {
			sum += treeList[i]-height
		}
	}
	return sum
}
