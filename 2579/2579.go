package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)
	pointList := make([]int, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscan(reader, &pointList[i])
	}

	maxPoint := make(map[int]int)
	maxPoint[0]=pointList[0]
	for i:=1 ; i<N ; i++ {
		max:=max(maxPoint[i-2],maxPoint[i-3]+pointList[i-1])
		maxPoint[i]=max + pointList[i]
	}
	fmt.Fprintln(writer, maxPoint[N-1])
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}