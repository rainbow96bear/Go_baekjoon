package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const MAX = 1000001

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, r, g, b int
	fmt.Fscan(reader, &n)

	cost := make([][3]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &r, &g, &b)
		cost[i][0] = r
		cost[i][1] = g
		cost[i][2] = b
	}

	ans := MAX

	for start := 0; start <= 2; start++ {
		dp := make([][3]int, n+1)

		for j := 0; j <= 2; j++ {
			if j == start {
				dp[1][start] = cost[1][start]
			} else {
				dp[1][j] = MAX
			}
		}

		for i := 2; i <= n; i++ {
			dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + cost[i][0]
			dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + cost[i][1]
			dp[i][2] = min(dp[i-1][1], dp[i-1][0]) + cost[i][2]
		}

		for j := 0; j <= 2; j++ {
			if j == start {
				continue
			}
			ans = int(math.Min(float64(ans), float64(dp[n][j])))
		}
	}

	fmt.Fprintln(writer, ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
