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

	list := []int{}
	for {
		var num int
		fmt.Fscanf(reader, "%d", &num)
		if num == 0 {
			break
		}
		list = append(list, num)
	}
	result := Process(list)
	fmt.Fprintln(writer, result)
}

func Process(list []int) int {
	dp := make([][][]int, len(list)+1)
	for i:=0 ; i<=len(list) ; i++ {
		dp[i] = make([][]int, 5)
		for j:=0 ; j<5  ; j++ {
			dp[i][j] = make([]int, 5)
			for k:=0 ; k<5 ; k++ {
				dp[i][j][k] = 1<<30
			}
		}
	}
	dp[0][0][0] = 0
	for i:=0 ; i<len(list) ; i++ {
		for L:=0 ; L<5 ; L++ {
			for R:=0 ; R<5 ; R++ {
				dp[i+1][list[i]][R] = min(dp[i+1][list[i]][R], dp[i][L][R] + CalcValue(L, list[i]))
				dp[i+1][L][list[i]] = min(dp[i+1][L][list[i]], dp[i][L][R] + CalcValue(R, list[i]))
			}
		}
	}

	minValue := 1<<30
	for L:=0 ; L<5 ; L++ {
		for R:=0 ; R<5 ; R++ {
			minValue = min(minValue, dp[len(list)][L][R])
		}
	}
	return minValue
}

func CalcValue(before, after int) int {
	if before == after  {
		return 1
	}
	if before == 0 {
		return 2
	}
	if abs(before - after) == 2 {
		return 4
	}
	return 3
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int) int {
	if a < b {
		 return a
	}
	return b
}