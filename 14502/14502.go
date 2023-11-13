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

	input := strings.Split(Input(scanner), " ")
	N, _ := strconv.Atoi(input[0])
	M, _ := strconv.Atoi(input[1])

	board := make([][]string, N)
	cnt := 0
	virusPositionList := []Position{} 
	for i:=0 ; i<N ; i++ {
		board[i] = strings.Split(Input(scanner), " ")
		for j:=0 ; j<M ; j++ {
			if board[i][j] == "2" {
				virusPositionList = append(virusPositionList, Position{i, j})
			}
			if board[i][j] == "0" {
				cnt++
			}
		}
	}

	result := Process(board, virusPositionList, N, M, cnt)
	fmt.Fprint(writer, result)
}

func Process(board [][]string, virusPositionList []Position, N, M, cnt int) int {
	max := 0
	
	for i:=0; i<N*M ; i++ {
		row1 := i/M
		col1 := i%M
		if board[row1][col1] == "0" {
			board[row1][col1] = "1"
			for j:=i+1 ; j<N*M ; j++ {
				row2 := j/M
				col2 := j%M
				if board[row2][col2] == "0" {
					board[row2][col2] = "1"

					for k:=j+1 ; k<N*M ; k++ {
						row3 := k/M
						col3 := k%M
						if board[row3][col3] == "0" {
							board[row3][col3] = "1"
							result, origin := BFS(board, virusPositionList, cnt-3, N, M)
							if max < result {
								max = result
							}
							board = origin
							board[row3][col3] = "0"							
						}
					}

					board[row2][col2] = "0"
				}
			}

			board[row1][col1] = "0"
		}
	}
	return max
}

func BFS(board [][]string, virusPositionList []Position, cnt, N, M int) (int, [][]string) {
	dx := []int{0, 0, 1, -1}
	dy := []int{-1, 1, 0, 0}
	count := cnt
	origin := make([][]string, N)
	for i:=0; i<N ; i++ {
		origin[i] = make([]string, M)
		for j:=0 ; j<M ; j++ {
			origin[i][j] = board[i][j]
		}
	}
	for len(virusPositionList) > 0 {
		now := virusPositionList[0]
		virusPositionList = virusPositionList[1:]
		for i:=0 ; i<4 ; i++ {
			nextX := now.x+dx[i]
			nextY := now.y+dy[i]
			if IsPossible(nextX, nextY, N, M) && board[nextX][nextY] == "0" {
				board[nextX][nextY] = "2"
				virusPositionList = append(virusPositionList, Position{nextX, nextY})
				count--
			}
		}
	}
	return count, origin
}

func IsPossible(x, y, N, M int) bool {
	if 0<=x && x<N && 0<=y && y<M {
		return true
	}
	return false
}

func Input(scanner *bufio.Scanner) string {
	scanner.Scan()

	return scanner.Text()
}