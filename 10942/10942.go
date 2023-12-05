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
	fmt.Fscanf(reader, "%d\n", &N)

	list := make([]int, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscanf(reader, "%d", &list[i])
	}
	fmt.Fscanln(reader)
	fmt.Fscanf(reader, "%d\n", &M)

	questionList := make([][]int, M)
	for i:=0 ; i<M ; i++ {
		var S, E int
		fmt.Fscanf(reader, "%d %d\n", &S, &E)
		questionList[i] = []int{S, E}
	}
	result := Process(N, M, list, questionList)
	for _, v := range result {
		fmt.Fprintln(writer, v)
	}
}

func Process(N, M int, list []int, questionList [][]int) []int {
	board := CheckPalindrome(N, list)
	result := []int{}
	for i:=0 ; i<M ; i++ {
		question := questionList[i]
		result = append(result, board[question[0]-1][question[1]-1])
	}
	return result
}

func CheckPalindrome(N int, list []int) [][]int {
	board := make([][]int, N)

	for i:=0; i<N ; i++ {
		board[i] = make([]int, N)
		board[i][i] = 1
		if i+1 < N {
			if list[i] == list[i+1] {
				board[i][i+1] = 1
			}else {
				board[i][i+1] = 0
			}
		}
	}
	for gap:=2 ; gap<N ; gap++ {
		for S:=0 ; S+gap<N ; S++ {
			if list[S] == list[S+gap] && board[S+1][S+gap-1] == 1{
				board[S][S+gap] = 1
			}else {
				board[S][S+gap] = 0
			}
		}
	}
	return board
}