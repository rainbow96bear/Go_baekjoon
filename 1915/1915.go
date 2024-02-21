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

	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	board := make([][]string, n)
	square := make([][]int, n+1)
	for i:=0 ; i<n ; i++ {
		board[i] = make([]string, m)
	}
	for i:=0 ; i<=n ; i++ {
		square[i] = make([]int, m+1)
	}
	for i:=0 ; i<n ; i++ {
		var input string
		fmt.Fscanf(reader, "%s\n", &input)
		for j, v := range input {
			board[i][j] = string(v)
		}
	}
	max := 0
	for i:=0 ; i<n ; i++ {
		for j:=0 ; j<m ; j++ {
			if board[i][j] == "1" {
				square[i+1][j+1] = min(square[i][j], min(square[i+1][j], square[i][j+1]))+1
				if max < square[i+1][j+1] {
					max = square[i+1][j+1]
				}
			}
		}
	}
	fmt.Fprintln(writer, max*max)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}