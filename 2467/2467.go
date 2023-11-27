package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscanln(reader, &N)
	list:=make([]int,N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscanf(reader, "%d", &list[i])
	}
	min := math.MaxInt
	answer := make([]int, 2)
	leftIdx, rightIdx := 0, N-1
	for leftIdx < rightIdx {
		sum := list[leftIdx]+list[rightIdx]
		if abs(sum) < abs(min) {
			min = sum
			answer[0]=list[leftIdx]
			answer[1]=list[rightIdx]
		}
		if sum == 0 {
			break	
		}
		if sum < 0 {
			leftIdx++
		}
		if sum > 0 {
			rightIdx--
		}
	}
	fmt.Fprintln(writer, answer[0], answer[1])
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}