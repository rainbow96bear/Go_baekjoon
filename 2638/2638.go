package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type Position struct {
	x int
	y int
}
func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()
	size := Input(scanner)
	N, M := size[0], size[1]
	
	board := make([][]int, N)
	for i:=0 ; i<N ; i++ {
		board[i] = Input(scanner)
	}
	result := Process(N, M, board)
	fmt.Fprintln(writer, result)
}

func Process(N, M int, board [][]int) int {
	needRepeat := true
	count := 0
	for needRepeat {
		count++
		needRepeat, board = FindSideCheese(N, M, board)
	}
	return count
}

func FindSideCheese(N, M int, board [][]int) (bool, [][]int) {
	list := []Position{{0, 0}}
	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}
	newBoard := make([][]int, N)

	isVisit := make([][]bool, N)
	for i:=0 ; i<N ; i++ {
		newBoard[i] = make([]int, M)
		isVisit[i] = make([]bool, M)
	}

	for len(list) > 0 {
		now := list[0]
		list = list[1:]

		for i:=0 ; i<4 ; i++ {
			nextX := now.x + dx[i]
			nextY := now.y + dy[i]

			if IsPossible(nextX, nextY, N, M) && isVisit[nextX][nextY] == false {
				if board[nextX][nextY] >= 1 {
					board[nextX][nextY]++
				}
				if board[nextX][nextY] == 0 {
					isVisit[nextX][nextY] = true
					list = append(list, Position{nextX,nextY})
				}
			}
		}
	}
	cnt := 0
	for i:=0 ; i<N ; i++ {
		for j:=0 ; j<M ; j++ {
			if board[i][j] == 1 || board[i][j] == 2 {
				newBoard[i][j] = 1
				cnt++
			}
		}
	}
	needRepeat := true
	if cnt == 0 {
		needRepeat = false
	}
	return needRepeat, newBoard
}

func IsPossible(x, y, N, M int) bool {
	if 0 <= x && x < N && 0 <= y && y < M {
		return true
	}
	return false
}

func Input(scanner *bufio.Scanner) []int {
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	numList := make([]int, len(input))
	for i, v := range input {
		num, _ := strconv.Atoi(v)
		numList[i] = num
	}
	return numList
}