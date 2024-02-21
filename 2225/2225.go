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
	fmt.Fscanf(reader, "%d %d\n", &N, &K)

	dp := make([][]int, K+1)
	for k:=0 ; k<=K ; k++ {
		dp[k]= make([]int, N+1)
	}
	for i:=0 ; i<=N ; i++ {
		dp[1][i] = 1
	}


	for k:=2 ; k<=K ; k++ {
		for n:=0 ; n<=N ; n++ {
			if n == 0 {
				dp[k][n] = dp[k-1][n]
			}else {
				dp[k][n] = (dp[k][n-1] + dp[k-1][n]) % 1000000000
			}
		}
	}
	fmt.Fprintln(writer, dp[K][N])
}