package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	board := make([][]int, 9)
	for i := 0; i < 9; i++ {
		board[i] = make([]int,9)
		Input(scanner, board[i])
	}

	Process(board, 0)

	for i := 0; i < 81; i++ {
		if i != 0 && i%9 == 0 {
			fmt.Fprintln(writer)
		}
		fmt.Fprintf(writer, "%d", board[i/9][i%9])
	}
}

func Process(board [][]int, n int) bool {
	for i := n; i < 81; i++ {
		x := i / 9
		y := i % 9
		if board[x][y] == 0 {
			for num := 1; num <= 9; num++ {
				if IsValid(board, x, y, num) {
					board[x][y] = num
					result := Process(board, i+1)
					if result {
						return true
					}
					board[x][y] = 0
				}
			}
			return false
		}
	}
	return true
}

func IsValid(board [][]int, x, y, num int) bool {
	for i := 0; i < 9; i++ {
		if board[i][y] == num || board[x][i] == num || board[(x/3)*3+i/3][(y/3)*3+i%3] == num {
			return false
		}
	}
	return true
}

func Input(scanner *bufio.Scanner, arr []int) {
	scanner.Scan()
	input := strings.Split(scanner.Text(), "")
	for i, v := range input {
		num, _ := strconv.Atoi(v)
		arr[i] = num
	}
}
