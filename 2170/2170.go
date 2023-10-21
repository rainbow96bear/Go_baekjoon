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

	var N int
	fmt.Fscan(reader, &N)

	arr := make([][]int,N)
	for i:=0 ; i<N ; i++ {
		arr[i] = make([]int, 2)
		fmt.Fscan(reader, &arr[i][0], &arr[i][1])
	}
	sort.Slice(arr, func(i, j int) bool {return arr[i][0]<arr[j][0]})
	total := 0
	left := arr[0][0]
	right := arr[0][1]
	for i:=1 ; i<N ; i++ {
		if arr[i][0] <= right {
			right = max(right, arr[i][1])
		}else if right < arr[i][0] {
			total += right-left
			left = arr[i][0]
			right = arr[i][1]
		}
	}
	total += right-left
	fmt.Fprintln(writer, total)
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}