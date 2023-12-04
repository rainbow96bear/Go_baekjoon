package main

import (
	"bufio"
	"fmt"
	"os"
)
type Node struct {
	length int
	result string
}
func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	firstString := Input(scanner)
	secondString := Input(scanner)

	length, result := Process(firstString, secondString)
	fmt.Fprintln(writer, length)
	fmt.Fprintln(writer, result)
}

func Process(firstString, secondeString string) (int, string) {

	board := make([][]Node, len(firstString)+1)

	for i:=0 ; i<=len(firstString) ; i++ {
		board[i] = make([]Node, len(secondeString)+1)
		board[i][0] = Node{0, ""}
	}

	for i:=0 ; i<=len(secondeString) ; i++ {
		board[0][i] = Node{0, ""}
	}

	for i:=1 ; i<=len(firstString) ; i++ {
		for j:=1 ; j<=len(secondeString) ; j++ {
			if firstString[i-1] == secondeString[j-1] {
				length := board[i-1][j-1].length + 1
				result := fmt.Sprint(board[i-1][j-1].result,string(secondeString[j-1]))
				board[i][j] = Node{length, result}
			}else {
				if board[i-1][j].length < board[i][j-1].length{
					board[i][j] =  board[i][j-1]
				}else {
					board[i][j] = board[i-1][j]
				}
			}
		}
	}

	return board[len(firstString)][len(secondeString)].length, board[len(firstString)][len(secondeString)].result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Input(scanner *bufio.Scanner) string {
	scanner.Scan()

	return scanner.Text()
}