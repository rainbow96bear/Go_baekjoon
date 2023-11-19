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
	brokenCount int
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	input := strings.Split(Input(scanner), " ")

	N, _ := strconv.Atoi(input[0])
	M, _ := strconv.Atoi(input[1])

	board := make([][]string, N)
	for i:=0 ; i<N ; i++ {
		board[i] = strings.Split(Input(scanner), "")
	}
	result := PathFind(N, M, board)
	fmt.Fprintln(writer, result)
}

func PathFind(N, M int, board [][]string) int {
	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}
	list := []Position{{0, 0, 0}}
	dpNoneBreak := make([][]int, N)
	dpBreak := make([][]int, N)
	for i:=0 ; i<N ; i++ {
		dpNoneBreak[i] = make([]int, M)
		dpBreak[i] = make([]int, M)
		for j:=0 ; j<M ; j++ {
			dpNoneBreak[i][j] = 1<<30
			dpBreak[i][j] = 1<<30
		}
	}
	dpNoneBreak[0][0] = 1
	dpBreak[0][0] = 1
	for len(list) > 0 {
		now := list[0]
		list = list[1:]

		for i:=0; i<4 ; i++ {
			nextX := now.x + dx[i]
			nextY := now.y + dy[i]
			if IsPossible(nextX, nextY, N, M) {
				if now.brokenCount == 1 {
					if board[nextX][nextY] == "0" {
						if dpBreak[nextX][nextY] > dpBreak[now.x][now.y]+1 {
							list = append(list, Position{nextX, nextY, now.brokenCount})
							dpBreak[nextX][nextY] = dpBreak[now.x][now.y]+1
						}
					}
				}
				if now.brokenCount == 0 {
					if board[nextX][nextY] == "0" {
						if dpNoneBreak[nextX][nextY] > dpNoneBreak[now.x][now.y]+1 {
							list = append(list, Position{nextX, nextY, now.brokenCount})
							dpNoneBreak[nextX][nextY] = dpNoneBreak[now.x][now.y]+1
						}
					}
					if board[nextX][nextY] == "1" {
						if now.brokenCount == 0 {
							if dpBreak[nextX][nextY] > dpNoneBreak[now.x][now.y]+1 {
								list = append(list, Position{nextX, nextY, now.brokenCount+1})
								dpBreak[nextX][nextY] = dpNoneBreak[now.x][now.y]+1
							}
						}
					}
				}
			}
		}
	}
	if dpBreak[N-1][M-1] == 1<<30 && dpNoneBreak[N-1][M-1] == 1<<30 {
		return -1
	}
	if dpBreak[N-1][M-1] > dpNoneBreak[N-1][M-1] {
		return dpNoneBreak[N-1][M-1]
	}
	return dpBreak[N-1][M-1]
}

func IsPossible(x, y, N, M int) bool {
	if 0 <= x && x < N && 0 <= y && y < M {
		return true
	}
	return false
}

func Input(scanner *bufio.Scanner) string {
	scanner.Scan()

	return scanner.Text()
}