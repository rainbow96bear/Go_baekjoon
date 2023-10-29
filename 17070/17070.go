package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pipe struct {
	state int
	r int
	c int
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	N, _ := strconv.Atoi(Input(scanner))

	board := make([][]string, N)
	for i:=0 ; i<N ; i++ {
		board[i] = strings.Split(Input(scanner), " ")
	}

	result := Process(board, N)

	fmt.Fprintln(writer, result)
}

func Process(board [][]string, N int) int {
	if board[N-1][N-1] == "1" {
		return 0
	}
	list := []Pipe{}
	list = append(list, Pipe{0, 0, 1})
	cnt := 0
	for len(list) > 0 {
		now := list[0]
		list = list[1:]
		if now.r == N-1 && now.c == N-1 {
			cnt++
			continue
		}
		if now.state == 0 {
			if checkHorizon(board, now, N) {
				list = append(list, Pipe{0, now.r, now.c+1})
			}
			if checkDiagonal(board, now, N) {
				list = append(list, Pipe{2, now.r+1, now.c+1})
			}
		}else if now.state == 1 {
			if checkVertical(board, now, N) {
				list = append(list, Pipe{1, now.r+1, now.c})
			}
			if checkDiagonal(board, now, N) {
				list = append(list, Pipe{2, now.r+1, now.c+1})
			}
		}else {
			if checkHorizon(board, now, N) {
				list = append(list, Pipe{0, now.r, now.c+1})
			}
			if checkVertical(board, now, N) {
				list = append(list, Pipe{1, now.r+1, now.c})
			}
			if checkDiagonal(board, now, N) {
				list = append(list, Pipe{2, now.r+1, now.c+1})
			}
		}
	}
	return cnt
}

func checkHorizon(board [][]string, pipe Pipe, N int) bool {
	if pipe.c+1 < N && board[pipe.r][pipe.c+1] != "1"{
		return true
	}
	return false
}

func checkVertical(board [][]string, pipe Pipe, N int) bool {
	if pipe.r+1 < N && board[pipe.r+1][pipe.c] != "1"{
		return true
	}
	return false
}
func checkDiagonal(board [][]string, pipe Pipe, N int) bool {
	if pipe.r+1 < N && pipe.c+1 < N && board[pipe.r+1][pipe.c+1] != "1" && checkVertical(board, pipe, N) && checkHorizon(board, pipe, N){
		return true
	}
	return false
}


func Input(scanner *bufio.Scanner) string {
	scanner.Scan()

	return scanner.Text()
}