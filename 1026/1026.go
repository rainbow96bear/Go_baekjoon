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

	var N int
	fmt.Fscan(reader, &N)

	A := make([]int, N)
	B := make([]int, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscan(reader, &A[i])
	}
	for i:=0 ; i<N ; i++ {
		fmt.Fscan(reader, &B[i])
	}

	sort.Ints(A)
	sort.Ints(B)
	sum := 0
	for i:=0 ; i<N ; i++{
		sum += A[i]*B[N-1-i]
	}
	fmt.Fprint(writer, sum)
}