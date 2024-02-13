package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()
	
	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	wine := make([]int, n+1)
	dp := make([]int, n+1)
	
	for i:=1 ; i<=n ; i++ {
		fmt.Fscanf(reader, "%d\n", &wine[i])
	}
	dp[0] = 0
	dp[1] = wine[1]


	if n > 1 {
		dp[2] = wine[1] + wine[2]
	}
	for i:=3 ; i<=n ; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+wine[i], dp[i-3]+wine[i]+wine[i-1])
	}

	fmt.Fprintln(writer, dp[n])
}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	}
	return c
}