package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()
	scanner.Scan()
	str1 := scanner.Text()
	scanner.Scan()
	str2 := scanner.Text()
	result := Process(str1, str2)
	fmt.Fprintln(writer, result)
}

func Process(str1, str2 string) int {
	board := make([][]int, len(str1)+1)
	for i:=0; i<=len(str1) ; i++ {
		board[i] = make([]int, len(str2)+1)
		for j:=0; j<=len(str2) ; j++ {
			if i == 0 || j == 0 {
				board[i][j] = 0
			}else if str1[i-1] == str2[j-1] {
				board[i][j] = board[i-1][j-1] + 1
			}else {
				board[i][j] = max(board[i-1][j], board[i][j-1])
			}
		}
	}
	return board[len(str1)][len(str2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}