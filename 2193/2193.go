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

	dp := make([][]int, N+1)

	for i:=1; i<=N ; i++ {
		dp[i] = make([]int, 2)
	}

	dp[1][0] = 0
	dp[1][1] = 1

	for i:=2 ; i<=N ; i++ {
		dp[i][0] = dp[i-1][0]+dp[i-1][1]
		dp[i][1] = dp[i-1][0]
	}

	fmt.Fprintln(writer,dp[N][0] + dp[N][1])
}