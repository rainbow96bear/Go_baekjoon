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

	K := 14
	N := 14
	board := make(map[int][]int)
	for i:=0 ; i<=K ; i++ {
		board[i] = append(board[i], 1)
		for j:=1 ; j<N ; j++ {
			if i==0 {
				board[i] = append(board[i], j+1)
			}else {
				board[i] = append(board[i], board[i-1][j]+board[i][j-1])
			}
		}
	}

	var T int
	fmt.Fscanln(reader, &T)

	for i:=0 ; i<T ; i++ {
		var k, n int
		fmt.Fscanln(reader, &k)
		fmt.Fscanln(reader, &n)
		fmt.Fprintln(writer, board[k][n-1])
	}
}