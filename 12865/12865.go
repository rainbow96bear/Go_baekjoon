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

	var N, K int
	fmt.Fscan(reader, &N, &K)

	weight := make([]int, N)
	value := make([]int, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscan(reader, &weight[i], &value[i])
	}
	result := Process(weight, value, N, K)
	fmt.Fprintln(writer, result)
}

func Process(weight, value []int , N, K int) int {
	dp := make(map[int]int)
	for i:=0 ; i<N ; i++ {
		for j:=K ; j>0 ; j-- {
			if weight[i] <= j {
				dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			}
		}
	}
	return dp[K]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}