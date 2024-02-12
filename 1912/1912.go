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
	fmt.Fscanf(reader, "%d\n", &n)
	list := make([]int, n)
	for i:=0 ; i<n ; i++ {
		fmt.Fscanf(reader, "%d", &list[i])
	}
	fmt.Fscanf(reader, "\n")

	dp := make([]int, n)
	answer := list[0]
	dp[0]=list[0]
	for i:=1 ; i<n ; i++ {
		dp[i] = max(dp[i-1]+list[i], list[i])
		if answer < dp[i] {
			answer = dp[i]
		}
	}

	fmt.Fprintln(writer, answer)
}

func max(a, b int) int {
	if a<b {
		return b
	}
	return a
}
