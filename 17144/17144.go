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

	inputRCT := Input(scanner)
	R, C, T := inputRCT[0], inputRCT[1], inputRCT[2]

	board := make([][]int, R)
	AirCleaner := []Position{}
	for i:=0 ; i<R ; i++ {
		board[i] = Input(scanner)
		for j:=0 ; j<C ; j++ {
			if board[i][j] == -1 {
				AirCleaner = append(AirCleaner, Position{i, j})
			}
		}
	}

	result := Process(R, C, T, board, AirCleaner)
	fmt.Fprintln(writer, result)
}

func Process(R, C, T int, board [][]int, AirCleaner []Position) int {
	for T > 0 {
		board = Spread(R, C, board)
		board = Turn(R, C, board, AirCleaner)
		T--
	}
	result := 0
	for i:=0 ; i<R ; i++ {
		for j:=0 ; j<C ; j++ {
			if board[i][j] > 0 {
				result += board[i][j]
			}
		}
	}
	return result
}

func Spread (R, C int, board [][]int) [][]int {
	newBoard := make([][]int, R)
	for i:=0 ; i<R ; i++ {
		newBoard[i] = make([]int, C)
	}
	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}
	for i:=0 ; i<R ; i++ {
		for j:=0 ; j<C ; j++ {
			if board[i][j] > 0 {
				newBoard[i][j] += board[i][j]
				for k:=0 ; k<4 ; k++ {
					nextX := i+dx[k]
					nextY := j+dy[k]
					if isPossible(nextX, nextY, R, C) && board[nextX][nextY] != -1 && board[i][j] > 0 {
						newBoard[nextX][nextY] += board[i][j]/5
						newBoard[i][j] -= board[i][j]/5
					}
				}
			}
		}
	}


	return newBoard
}

func Turn(R, C int, board [][]int, AirCleaner []Position) [][]int {
	up := AirCleaner[0]
	down := AirCleaner[1]
	board[up.x][up.y] = 0
	board[down.x][down.y] = 0
	newUpTop := make([]int, len(board[0][1:]))
	copy(newUpTop, board[0][1:])
	newUpBottom := make([]int, len(board[up.x][:C-1]))
	copy(newUpBottom, board[up.x][:C-1])
	for i:=0 ; i<up.x ; i++ {
		board[i][C-1] = board[i+1][C-1]
		board[up.x-i][0] = board[up.x-i-1][0]
	}
	for i:=0 ; i<len(newUpTop) ; i++ {
		board[0][i] = newUpTop[i]
		board[up.x][i+1] = newUpBottom[i]
	}
	newDownTop := make([]int, len(board[down.x][:C-1]))
	copy(newDownTop, board[down.x][:C-1])
	newDownBottom := make([]int, len(board[R-1][1:]))
	copy(newDownBottom, board[R-1][1:])
	for i:=down.x+1 ; i<R ; i++ {
		board[i-1][0] = board[i][0]
		board[R+down.x-i][C-1] = board[R+down.x-1-i][C-1]
	}
	for i:=0 ; i<len(newDownTop) ; i++ {
		board[R-1][i] = newDownBottom[i]
		board[down.x][i+1] = newDownTop[i]
	}
	board[up.x][up.y] = -1
	board[down.x][down.y] = -1
	return board
}

func isPossible(x, y, R, C int) bool {
	if 0 <= x && x < R && 0 <= y && y < C {
		return true
	}
	return false
}

func Input(scanner *bufio.Scanner) []int {
	scanner.Scan()

	input := scanner.Text()

	inputList := strings.Split(input, " ")

	result := []int{}
	for _, v := range inputList {
		num, _ := strconv.Atoi(v)
		result = append(result, num)
	}
	return result
}