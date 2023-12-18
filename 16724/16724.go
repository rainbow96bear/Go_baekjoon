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

	N, M := InputNum(scanner)
	board := make([][]string, N)
	for i:=0 ; i<N ; i++ {
		board[i] = InputBoard(scanner)
	}
	
	result := FindMinNumOfSafeZone(N, M, board)

	fmt.Fprintln(writer, result)
}

func InputNum(scanner *bufio.Scanner) (int,int){
	scanner.Scan()

	input := strings.Split(scanner.Text(), " ")

	N, _ := strconv.Atoi(input[0])
	M, _ := strconv.Atoi(input[1])

	return N, M
}

func InputBoard(scanner *bufio.Scanner) []string {
	scanner.Scan()

	input := strings.Split(scanner.Text(), "")

	return input
}

func FindMinNumOfSafeZone(N, M int, board [][]string) int {
	isVisit := make([][]bool, N)
	Root := make([][]int, N)
	for i:=0 ; i<N ; i++ {
		isVisit[i] = make([]bool, M)
		Root[i] = make([]int, M)
		for j:=0 ; j<M ; j++ {
			Root[i][j] = i*M+j
		}
	}
	
	for i:=0 ; i<N ; i++ {
		for j:=0 ; j<M ; j++ {
			if isVisit[i][j] == false {
				var root int
				root = FindRoot(i, j, board, Root, isVisit)
				Root[i][j] = root
			}
		}
	}
	check := make([]bool,N*M)
	count :=0
	for i:=0 ; i<N ; i++ {
		for j:=0 ; j<M ; j++ {
			if check[Root[i][j]] == false {
				check[Root[i][j]] = true
				count++
			}
		}
	}
	return count
}

func FindRoot(row, col int, board [][]string, Root [][]int, isVisit [][]bool) int {
	if isVisit[row][col] {
		return Root[row][col]
	}
	isVisit[row][col] = true
	var root int
	switch board[row][col] {
	case "U" :
		root = FindRoot(row-1, col, board, Root, isVisit)
	case "D" :
		root = FindRoot(row+1, col, board, Root, isVisit)
	case "L" :
		root = FindRoot(row, col-1, board, Root, isVisit)
	case "R" :
		root = FindRoot(row, col+1, board, Root, isVisit)
	}
	Root[row][col] = root
	return root
}