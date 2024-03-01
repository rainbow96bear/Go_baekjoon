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

	var N, M int
	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	board := make([][]int, N+1)
	board[0] = make([]int, M+1)
	for i:=1 ; i<=N ; i++ {
		board[i] = make([]int, M+1)
		for j:=1 ; j<=M ; j++ {
			var input int
			fmt.Fscanf(reader, "%d", &input)
			board[i][j] = board[i][j-1]+input
		}
		fmt.Fscanf(reader, "\n")
	}

	var K int
	fmt.Fscanf(reader, "%d\n", &K)

	for i:=0 ; i<K ; i++ {
		var i, j, x, y int
		fmt.Fscanf(reader, "%d %d %d %d\n", &i, &j, &x, &y)

		total := 0
		for index:=i ; index<=x ; index++ {
			total += (board[index][y]-board[index][j-1])
		}
		fmt.Fprintln(writer, total)
	}
}