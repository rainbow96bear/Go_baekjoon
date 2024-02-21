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

	var n, k int
	fmt.Fscanf(reader, "%d %d\n", &n, &k)

	coinList := make([]int, n)

	for i:=0 ; i<n ; i++ {
		fmt.Fscanf(reader, "%d\n", &coinList[i])
	}

	dp := make(map[int]int)
	dp[0] = 1
	for _, v := range coinList {
		for i:=v ; i<=k ; i++ {
			dp[i] += dp[i-v]
		}
	}
	fmt.Fprintln(writer, dp[k])
}