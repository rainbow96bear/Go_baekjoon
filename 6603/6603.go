package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	N       int
	numbers []int
	result  []int
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	

	defer writer.Flush()
	
	for {
		fmt.Fscanf(reader, "%d", &N)
		if N == 0 {
			return
		}

		numbers = make([]int, N)
		result = make([]int, 6)

		for i := 0; i < N; i++ {
			fmt.Fscanf(reader, "%d", &numbers[i])
		}
		fmt.Fscanf(reader, "\n")
		backtrack(0, 0)
		fmt.Fprintf(writer, "\n")
	}
}

func backtrack(start, depth int) {
	if depth == 6 {
		for _, v := range result {
			fmt.Fprintf(writer, "%d ", v)
		}
		fmt.Fprintln(writer)
		return
	}

	for i := start; i < N; i++ {
		result[depth] = numbers[i]
		backtrack(i+1, depth+1)
	}
}