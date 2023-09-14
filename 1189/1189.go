package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var R, C, K int
	fmt.Fscan(reader, &R, &C, &K)
	board := make([][]string,R)
	for i:=0 ; i<R ; i++ {
		var input string
		fmt.Fscan(reader, &input)
		board[i] = strings.Split(input,"")
	}
	fmt.Fprint(writer, PathFind(board, R-1, 0, R, C, K-1))
}

func PathFind(board [][]string, r, c, R, C, K int) int {
	if K == 0 {
		if r == 0 && c == C-1 {
			return 1
		}else {
			return 0
		}
	}
	original := board[r][c]
	board[r][c] = "T"
	answer := 0 
	if 0 <= r-1 && board[r-1][c] != "T"{
		answer += PathFind(board, r-1, c, R, C, K-1)
	}
	if r+1 < R && board[r+1][c] != "T"{
		answer += PathFind(board, r+1, c, R, C, K-1)
	}
	if 0 <= c-1 && board[r][c-1] != "T"{
		answer += PathFind(board, r, c-1, R, C, K-1)
	}
	if c+1 < C && board[r][c+1] != "T"{
		answer += PathFind(board, r, c+1, R, C, K-1)
	}
	board[r][c] = original
	return answer
}

