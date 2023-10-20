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
	fmt.Fscan(reader, &N, &M)

	arr := make([][]int, N)
	for i:=0 ; i<N ; i++ {
		arr[i] = make([]int, M)
		for j:=0 ; j<M ; j++ {
			fmt.Fscan(reader, &arr[i][j])
		}
	}
	dp := make([][][]int, N)
	for i:=0 ; i<N ; i++ {
		dp[i] = make([][]int, M)
	}
	for i:=0 ; i<M ; i++ {
		dp[0][i] = make([]int, 3)
		dp[0][i][0] = arr[0][i]
		dp[0][i][1] = arr[0][i]
		dp[0][i][2] = arr[0][i]
	}
	for i:=1 ; i<N ; i++ {
		for j:=0 ; j<M ; j++ {
			dp[i][j] = make([]int, 3)
			if j == 0 {
				dp[i][j][1] = dp[i-1][j][2] + arr[i][j]
				min := min(dp[i-1][j+1][0], dp[i-1][j+1][1])
				dp[i][j][2] = min + arr[i][j]
			}else if j == M-1 {
				dp[i][j][1] = dp[i-1][j][0] + arr[i][j]
				min := min(dp[i-1][j-1][1], dp[i-1][j-1][2])
				dp[i][j][0] = min + arr[i][j]
			}else {
				min0 := min(dp[i-1][j-1][1], dp[i-1][j-1][2])
				dp[i][j][0] = min0 + arr[i][j]
				min1 := min(dp[i-1][j][0], dp[i-1][j][2])
				dp[i][j][1] = min1 + arr[i][j]
				min2 := min(dp[i-1][j+1][0], dp[i-1][j+1][1])
				dp[i][j][2] = min2 + arr[i][j]
			}
		}
	}
	minValue := min(dp[N-1][0][1], dp[N-1][0][2])

	for i:=1 ; i<M-1 ; i++ {
		for j:=0 ; j<3 ; j++ {
			minValue = min(minValue, dp[N-1][i][j])
		}
	}
	minValue = min(minValue, min(dp[N-1][M-1][0], dp[N-1][M-1][1]))
	fmt.Fprintln(writer, minValue)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
