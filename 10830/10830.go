package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	input := StrInput(scanner)
	N, _ := strconv.Atoi(input[0])
	B, _ := strconv.ParseInt(input[1], 10, 64)

	board := make([][]int, N)
	for i, _ := range board {
		board[i] = NumInput(scanner)
	}

	result := Process(B, N, board)

	for _, v := range result {
		for _, answer := range v {
			fmt.Fprint(writer, answer, " ")
		}
		fmt.Fprintln(writer)
	}
}

func Process(B int64, N int, board [][]int) [][]int {
	answer := make([][]int, N)
	for i:=0 ; i<N ; i++ {
		answer[i] = make([]int, N)
		answer[i][i] = 1
	}
	
	for B > 0 {
		if B%2 == 1 {
			answer = calc(answer, board, N)
		}
		board = calc(board, board, N)
		B /= 2
	}
	return answer
}

func calc(base, board [][]int, N int) [][]int {
	result := make([][]int,N)
	for i:=0 ; i<N ; i++ {
		result[i] = make([]int, N)
		for j:=0; j<N ; j++ {
			for k:=0 ; k<N ; k++ {	
				result[i][j] += (base[i][k] * board[k][j])%1000
			}
			result[i][j] %= 1000
		}
	}
	return result
}

func StrInput(scanner *bufio.Scanner) []string {
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")

	return input
}

func NumInput(scanner *bufio.Scanner) []int {
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	list := []int{}
	for _, v := range input {
		num, _ := strconv.Atoi(v)
		list = append(list, num)
	}
	return list
}