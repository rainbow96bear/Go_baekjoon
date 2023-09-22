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

	dp := make([]int, n+1)

    dp[1] = 1
    dp[2] = 2

    for i := 3; i <= n; i++ {
        dp[i] = (dp[i-1] + dp[i-2]) % 10007
    }

	fmt.Fprintln(writer, dp[n])
}