package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N, M int
	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	left := []int{}
	right := []int{}
	leftMax := 0
	rightMax := 0
	for i:=0 ; i<N ; i++ {
		var input int
		fmt.Fscanf(reader, "%d ", &input)
		if input < 0 {
			left = append(left, -input)
			if leftMax < -input {
				leftMax = -input
			}
		}else {
			right = append(right, input)
			if rightMax < input {
				rightMax = input
			}
		}
	}
	sortToDown(left)
	sortToDown(right)
	leftIndex := 0
	rightIndex := 0
	answer := 0
	
	for ; leftIndex<len(left) && rightIndex<len(right) ;{
		if left[leftIndex] > right[rightIndex] {
			answer += 2*left[leftIndex]
			leftIndex+=M
		}else {
			answer += 2*right[rightIndex]
			rightIndex+=M
		}
	}
	for ;leftIndex < len(left) ; {
		answer += 2*left[leftIndex]
		leftIndex+=M
	}
	for ; rightIndex < len(right) ; {
		answer += 2*right[rightIndex]
		rightIndex+=M
	}
	if rightMax < leftMax {
		answer -= leftMax
	}else {
		answer -= rightMax
	}
	fmt.Fprintln(writer, answer)
}

func sortToDown(slice []int) {
	sort.Slice(slice, func(i, j int)bool{return slice[i]>slice[j]})
}