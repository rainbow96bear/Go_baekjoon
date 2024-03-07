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

	var N, M int
	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	boxes := make([]int, N)
	books := make([]int, M)
	for i:=0 ; i<N ; i++ {
		fmt.Fscanf(reader, "%d", &boxes[i])
	}
	fmt.Fscanf(reader, "\n")

	for  i:=0 ; i<M ; i++ {
		fmt.Fscanf(reader, "%d", &books[i])
	}
	fmt.Fscanf(reader, "\n")

	garbage, bookIndex := 0, 0
	for i:=0 ; i<N ; i++ {
		total := 0
		for bookIndex<M && total + books[bookIndex] <= boxes[i] {
			total += books[bookIndex]
			bookIndex++
		}
		garbage += (boxes[i]-total)
	}

	fmt.Fprintln(writer, garbage)
}