package main

import (
	"bufio"
	"fmt"
	"os"
)
var (
	divide int = 1000000000
)
func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()
	
	var N int
	fmt.Fscanf(reader, "%d\n", &N)

	dp := make([][]int, N+1)
	for i:= range dp {
		dp[i] = make([]int, 10)
	}

	for i:=1 ; i<10 ; i++ {
		dp[1][i] = 1
	}

	for i:=2 ; i<=N ; i++ {
		for j:=0 ; j<10 ; j++ {
			if j-1 >= 0 {
				dp[i][j] = (dp[i][j]+dp[i-1][j-1])%divide
			}
			if j+1 <10 {
				dp[i][j] = (dp[i][j]+dp[i-1][j+1])%divide
			}
		}
	}
	answer := 0
	for  i := 0; i <= 9; i++ {
		answer = (answer + dp[N][i])%divide
	}
	fmt.Fprintln(writer, answer)
}