package main

import (
	"bufio"
	"fmt"
	"os"
)
var N, P, Q int
var dp map[int]int

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	fmt.Fscanf(reader, "%d %d %d\n", &N, &P, &Q)
	dp = make(map[int]int)
	dp[0] = 1
	answer := calc(N)
	
	fmt.Fprintln(writer, answer)
}

func calc(Num int) int {
	if dp[Num] == 0 {
		dp[Num] = calc(Num/P) + calc(Num/Q)
	}
	return dp[Num]
}