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
	fmt.Fscanf(reader, "%d\n", &N)

	list := make([]int, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscanf(reader, "%d ", &list[i])
	}
	fmt.Fscanf(reader, "\n")

	result := FindMaxLength(N, list)
	fmt.Fprintln(writer, result)
}

func FindMaxLength(N int, list []int) int {
	part := []int{list[0]}
	for i:=0 ; i<N ; i++ {
		now := list[i]
		if part[len(part)-1] < now {
			part = append(part, now)
		}else {
			start := 0
			end := len(part)-1
			for start < end {
				mid := (start + end)/2
				if part[mid] < now {
					start = mid +1
				}else{
					end = mid
				}
			}
			part[end] = now
		}
	}
	return len(part)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}