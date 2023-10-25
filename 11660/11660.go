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
	fmt.Fscan(reader, &N, &M)

	board := make([][]int, N)
	for i:=0 ; i<N ; i++ {
		board[i] = make([]int, N)
		for j:=0 ; j<N ; j++ {
			fmt.Fscan(reader, &board[i][j])
		}
	}
	testCase := make([][]int, M)
	for i:=0 ; i<M ; i++ {
		var fromX, fromY, toX, toY int
		fmt.Fscan(reader, &fromX, &fromY, &toX, &toY)
		testCase[i] = []int{fromX, fromY, toX, toY}
	}
	result := Process(board, testCase, N)

	for _, v := range result {
		fmt.Fprintln(writer, v)
	}
}

func Process(board, testCase [][]int, N int) []int {
	sumBoard := make([][]int, N+1)
	
	for i:=0 ; i<=N ; i++ {
		sumBoard[i] = make([]int, N+1)
		for j:=0 ; j<=N ; j++ {
			if i == 0 || j == 0 {
				sumBoard[i][j] = 0
			} else {
				sumBoard[i][j] = sumBoard[i-1][j] + sumBoard[i][j-1] - sumBoard[i-1][j-1] + board[i-1][j-1]
			}
		}
	}
	
	result := []int{}
	for _, v := range testCase {
		total := sumBoard[v[2]][v[3]] + sumBoard[v[0]-1][v[1]-1] - sumBoard[v[2]][v[1]-1] - sumBoard[v[0]-1][v[3]]
		result = append(result, total)
	}
	return result
}