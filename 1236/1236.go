package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var r, c int
	fmt.Fscanf(reader, "%d %d\n", &r, &c)

	board := make([][]string, r)
	for i:=0 ; i<r ; i++ {
		board[i] = make([]string, c)
		var row string
		fmt.Fscanf(reader, "%s\n", &row)
		board[i] = strings.Split(row, "")
	}

	result := findMin(r, c, board)
	fmt.Fprintln(writer, result)

}

func findMin(r, c int, board [][]string) int {
	row := r
	for i:=0 ; i<r ; i++ {
		for j:=0 ; j<c ; j++ {
			if board[i][j] == "X" {
				row--
				break
			}
		}
	}
	col := c
	for i:=0 ; i<c ; i++ {
		for j:=0 ; j<r ; j++ {
			if board[j][i] == "X" {
				col--
				break
			}
		}
	}
	if row > col {
		return row
	}
	return col
}