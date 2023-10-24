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

	var T int
	fmt.Fscan(reader, &T)

	for i:=0 ; i<T ; i++ {
		var n int
		fmt.Fscan(reader, &n)
		sticker := make([][]int,2)
		sticker[0] = make([]int, n)
		sticker[1] = make([]int, n)
		for j:=0 ; j<n ; j++ {
			fmt.Fscan(reader, &sticker[0][j])
		}
		for j:=0 ; j<n ; j++ {
			fmt.Fscan(reader, &sticker[1][j])
		}
		answer := Process(sticker, n)
		fmt.Fprintln(writer, answer)
	}
}

func Process(sticker [][]int, n int) int {
	dp := make(map[int]map[int]int)
	dp[0] = make(map[int]int)
	dp[1] = make(map[int]int)
	dp[0][0] = sticker[0][0]
	dp[1][0] = sticker[1][0]
	for i:=0 ; i<n ; i++ {
		dp[0][i] =max(dp[1][i-1] + sticker[0][i], dp[1][i-2] + sticker[0][i])
		dp[1][i] =max(dp[0][i-1] + sticker[1][i], dp[0][i-2] + sticker[1][i])
	}
	return max(dp[0][n-1], dp[1][n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}