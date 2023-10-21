package main

import (
	"bufio"
	"fmt"
	"os"
)
const mod = 1000000
func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)
	if N == 1 {
		fmt.Fprintln(writer, 3)
		return
	}
	dp := make([][][]int, N+1)
	// dp[총 일수][지각 횟수][A의 연속 개수]
	for i:=0 ; i<=N ; i++ {
		dp[i] = make([][]int,2)
		dp[i][0] = make([]int, 3)
		dp[i][1] = make([]int, 3)
	}

	dp[1][0][0] = 1
	dp[1][0][1] = 1
	dp[1][0][2] = 0
	dp[1][1][0] = 1
	dp[1][1][1] = 0
	dp[1][1][2] = 0

	for i:=2 ; i<=N ; i++ {
		dp[i][0][0] = (dp[i-1][0][0] + dp[i-1][0][1] + dp[i-1][0][2])%1000000
		dp[i][0][1] = dp[i-1][0][0]
		dp[i][0][2] = dp[i-1][0][1]
		dp[i][1][0] = (dp[i-1][0][0] + dp[i-1][0][1] + dp[i-1][0][2] + dp[i-1][1][0] + dp[i-1][1][1] + dp[i-1][1][2])%1000000
		dp[i][1][1] = dp[i-1][1][0]
		dp[i][1][2] = dp[i-1][1][1]
	}
	result := (dp[N][0][0] + dp[N][0][1] + dp[N][0][2] + dp[N][1][0] + dp[N][1][1] + dp[N][1][2])%1000000
	fmt.Fprintln(writer, result)
}