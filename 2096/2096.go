package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)

	maxDP := make([]int, 3) 
	minDP := make([]int, 3)
	var A, B, C int
	fmt.Fscan(reader, &A, &B, &C)
	maxDP[0] = A
	maxDP[1] = B
	maxDP[2] = C
	minDP[0] = A
	minDP[1] = B
	minDP[2] = C
	for i:=1 ; i<N ; i++ {
		var a, b, c int
		fmt.Fscan(reader, &a, &b, &c)

		newMaxDP := make([]int, 3) 
		newMinDP := make([]int, 3)

		newMaxDP[0] = findMax(maxDP[0], maxDP[1]) + a
		newMaxDP[1] = findMax(findMax(maxDP[0], maxDP[1]), maxDP[2]) + b
		newMaxDP[2] = findMax(maxDP[1], maxDP[2]) + c

		newMinDP[0] = findMin(minDP[0], minDP[1]) + a
		newMinDP[1] = findMin(findMin(minDP[0], minDP[1]), minDP[2]) + b
		newMinDP[2] = findMin(minDP[1], minDP[2]) + c
		copy(maxDP, newMaxDP)
		copy(minDP, newMinDP)
	}
	max := findMax(findMax(maxDP[0], maxDP[1]), maxDP[2])
	min := findMin(findMin(minDP[0], minDP[1]), minDP[2])
	fmt.Fprintln(writer, max, min)
}

func findMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func findMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}