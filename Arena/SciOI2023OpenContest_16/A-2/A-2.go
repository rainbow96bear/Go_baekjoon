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

	var N, K int
	fmt.Fscanf(reader, "%d %d\n", &N, &K)

	Alist := make([]int, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscanf(reader, "%d ", &Alist[i])
	}
	fmt.Fscanf(reader, "\n")
	result := FindMinCount_Add(N, K, Alist)

	fmt.Fprintln(writer, result)
}

func FindMinCount_Add(N, K int, Alist []int) (count int) {
	count = 0
	for i:=1 ; i<N ; i++ {
		if Alist[i-1] >= Alist[i] {
			if Alist[i-1] < Alist[i]+K {
				Alist[i] += K
				count++
			}else {
				return -1
			}
		}
	}
	return count
}