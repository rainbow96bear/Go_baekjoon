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
	fmt.Fscanf(reader, "%d\n", &N)

	list := make([]int, N)

	for i:=0 ; i<N ; i++ {
		fmt.Fscanf(reader, "%d", &list[i])
	}
	fmt.Fscanf(reader, "\n")
	dp := make([][2]int, N)

	dp[0][0] = list[0]
	dp[0][1] = list[0]
	maxSum := list[0]

	for i := 1; i < N; i++ {
		dp[i][0] = max(dp[i-1][0]+list[i], list[i])
		dp[i][1] = max(dp[i-1][0], dp[i-1][1]+list[i])

		maxSum = max(maxSum, max(dp[i][0], dp[i][1]))
	}

	fmt.Println(maxSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}