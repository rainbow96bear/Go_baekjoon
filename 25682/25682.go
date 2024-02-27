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

	var N, M, K int
	fmt.Fscanf(reader, "%d %d %d\n", &N, &M, &K)

	min, max := K*K, 0

	countBoard := make([][]int, N)
	for i:=0 ; i<N ; i++ {
		var input string
		fmt.Fscanf(reader, "%s\n", &input)

		countBoard[i] = make([]int, M)
		
		for j, v := range input {
			if !(((i+j)%2 == 0 && v == 'B') || ((i+j)%2 == 1 && v == 'W')){
				countBoard[i][j]++
			}
		}
	}
	
	dp := make([][]int, N+1)
	dp[0] = make([]int, M+1)
	for i:=1 ; i<=N ; i++ {
		dp[i] = make([]int, M+1)
		for j:=1 ; j<=M ; j++ {
			dp[i][j] = countBoard[i-1][j-1] - dp[i-1][j-1] + dp[i][j-1] + dp[i-1][j]
		}
	}

	for i:=K ; i<=N ; i++ {
		for j:=K ; j<=M ; j++ {
			calc := dp[i][j] + dp[i-K][j-K] - dp[i-K][j] - dp[i][j-K]
			if min > calc {
				min = calc
			}
			if max < calc {
				max = calc
			}
		}
	}

	if min > K*K - max {
		fmt.Fprintln(writer, K*K-max)
	}else {
		fmt.Fprintln(writer, min)
	}
}