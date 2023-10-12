package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	x int
	y int
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	board := make([][]int, n)
	isVisit := make([][]bool, n)
	for i:=0 ; i<n ; i++ {
		board[i] = make([]int, m)
		isVisit[i] = make([]bool, m)
		for j:=0 ; j<m ; j++ {
			fmt.Fscan(reader, &board[i][j])
		}
	}
	max := 0
	for i:=0 ; i<n ; i++ {
		for j:=0 ; j<m ; j++ {
			isVisit[i][j] = true
			result := DFS(board, isVisit, n, m , i, j, 1)
			isVisit[i][j] = false
			if max < result {
				max = result
			}
		}
	}
	fmt.Fprintln(writer, max)
}

func DFS(board [][]int, isVisit [][]bool, n, m, x, y, cnt int) int {
	max := 0
	if cnt == 4 {
		return board[x][y]
	}
	if cnt == 2 {
		move := [][]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}
		for i:=0 ; i<4 ; i++ {
			X := x + move[i][0]
			Y := y + move[i][1]

			if 0 <= X && X < n && 0 <= Y && Y < m {
				if isVisit[X][y] == false && isVisit[x][Y] == false {
					result := board[X][y] + board[x][Y]
					if max < result {
						max = result
					}
				}
			}
		}
	}

	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}
	for i:=0 ; i<4 ; i++ {
		X := x + dx[i]
		Y := y + dy[i]

		if 0 <= X && X < n && 0 <= Y && Y < m {
			if isVisit[X][Y] == false {
				isVisit[X][Y] = true
				result := DFS(board, isVisit, n, m , X, Y, cnt+1)
				if max < result {
					max = result
				}
				isVisit[X][Y] = false
			}
		}
	}
	return max + board[x][y]
}