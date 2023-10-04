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

	var N int
	fmt.Fscan(reader, &N)

	board := make([][]int, N)
	for i:=0 ; i<N ; i++ {
		for j:=0 ; j<N ; j++ {
			var input int
			fmt.Fscan(reader, &input)
			board[i] = append(board[i], input)
		}
	}
	check := make([][]int,N)
	for i:=0 ; i<N ; i++ {
		queue := []int{i}
		isVisit := make([]bool, N)
		check[i] = make([]int, N)
		for len(queue) > 0 {
			now := queue[0]
			queue = queue[1:]
			isVisit[now] = true
			for j:=0 ; j<N ; j++ {
				if i == j && board[now][j] == 1{
					check[i][i] = 1
				}
				if board[now][j] == 1 && isVisit[j] == false {
					check[i][j] = 1
					queue = append(queue, j)
				}
			}
		}
	}

	for i:=0 ; i<N ; i++ {
		for j:=0 ; j<N ; j++ {
			fmt.Fprintf(writer, "%d ", check[i][j])
		}
		fmt.Fprintln(writer)
	}
}