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

	var n int
	fmt.Fscan(reader, &n)

	arr := make([][]int, n)
	for i:=0 ; i<n ; i++ {
		arr[i] = make([]int, i+1)
		for j:=0 ; j<i+1 ; j++ {
			fmt.Fscan(reader, &arr[i][j])
		}
	}
	answer := Process(n, arr)
	fmt.Fprintln(writer, answer)
}

func Process(n int, arr [][]int) int {
	dp := make([][]int, n)
	for i:=0 ; i<n ; i++ {
		dp[i] = make([]int, i+1)
	}
	dp[0] = []int{arr[0][0]}

	for i:=1 ; i<n ; i++ {
		for j:=0 ; j<=i ; j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + arr[i][j]
			}else if j == i {
				dp[i][j] = dp[i-1][j-1] + arr[i][j]
			}else {
				dp[i][j] = max(dp[i-1][j] + arr[i][j], dp[i-1][j-1] + arr[i][j])
			}
		}
	}
	answer := 0
	for i:=0 ; i<n ; i++ {
		answer = max(answer, dp[n-1][i])
	}
	return answer
}
func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}