package main

import (
	"bufio"
	"fmt"
	"os"
)

type Matrix struct {
	row int
	col int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscanf(reader, "%d\n", &N)

	MatrixList := make([]Matrix, N)

	for i:=0 ; i<N ; i++ {
		var r, c int
		fmt.Fscanf(reader, "%d %d\n", &r, &c)
		MatrixList[i] = Matrix{r, c}
	}

	result := FindMinCalc(N, MatrixList)
	fmt.Fprintln(writer, result)
}


func FindMinCalc(N int, MatrixList []Matrix) int {

	dp := make([][]int, N)

	for i:=0 ; i<N ; i++{
		dp[i] = make([]int, N)
	}

	for term:=1 ; term<N ; term++ {
		for start:=0 ; start<N ; start++ {
			if start+term == N {
				break
			}
			dp[start][start+term] = 1<<30
			for t:=start ; t<start+term ; t++ {
				dp[start][start+term] = min(dp[start][start+term], dp[start][t]+dp[t+1][start+term]+MatrixList[start].row*MatrixList[t].col*MatrixList[start+term].col)
			}
		}
	}

	return dp[0][N-1]
}