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

	for i:=0 ; i<N ; i++ {
		board[i] = strings.Split(Input(scanner), " ")
	}
	answer := Process(board, N, M)
	fmt.Fprintln(writer, answer)
}

func Process(board [][]string, N, M int) int {
	homeList := []Position{}
	shopList := []Position{}
	for i:=0 ; i<N ; i++ {
		for j:=0 ; j<N ; j++ {
			if board[i][j] == "1" {
				homeList = append(homeList, Position{i, j})
			}else if board[i][j] == "2" {
				shopList = append(shopList, Position{i, j})
			}
		}
	}
	selectList := []Position{}
	result := calc(shopList, homeList, selectList, M)
	return result
}

func calc(shopList, homeList, selectList []Position, M int) int {
	if M == 0 {
		
		total := 0
		for i:=0 ; i<len(homeList) ; i++ {
			min := abs(homeList[i].x - selectList[0].x) + abs(homeList[i].y - selectList[0].y)
			for j:=1 ; j<len(selectList) ; j++ {
				length := abs(homeList[i].x - selectList[j].x) + abs(homeList[i].y - selectList[j].y)
				if min > length {
					min = length
				}
			}
			total += min
		}
		return total
	}
	if len(shopList) == 0 {
		return -1
	}
	nextSelectList := []Position{}
	nextSelectList = append(nextSelectList, selectList...)
	nextSelectList = append(nextSelectList, shopList[0])
	result1 := calc(shopList[1:], homeList, nextSelectList, M-1)
	result2 := calc(shopList[1:], homeList, selectList, M)
	if result1 == -1 {
		return result2
	}
	if result2 == -1 {
		return result1
	}
	if result1 > result2 {
		return result2
	}
	return result1
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Input(scanner *bufio.Scanner) string {
	scanner.Scan()

	input := scanner.Text()

	return input
}