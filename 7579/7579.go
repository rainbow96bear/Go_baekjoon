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

	var N, M int
	fmt.Fscanf(reader, "%d %d\n", &N, &M)
	MemoryList := make([]int, N)
	CostList := make([]int, N)

	for i:=0; i< N ; i++ {
		fmt.Fscanf(reader, "%d", &MemoryList[i])
	}
	fmt.Fscanf(reader, "\n")

	for i:=0 ; i<N ; i++ {
		fmt.Fscanf(reader, "%d", &CostList[i])
	}
	fmt.Fscanf(reader, "\n")

	result := FindMinCost(N, M, MemoryList, CostList)
	fmt.Fprintln(writer, result)
}


func FindMinCost(N, M int, MemoryList, CostList []int) int {
	MaxCost := 10000
	dp := make([]int, MaxCost+1)

	for i:=0 ; i<N ; i++ {
		for j:=MaxCost ; j>=CostList[i] ; j-- {
			dp[j] = max(dp[j], dp[j-CostList[i]]+MemoryList[i])
		}
	}
	var answer int
	for i:=0 ; i<=MaxCost ; i++ {
		if dp[i] >= M {
			answer = i
			break
		}
	}
	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}