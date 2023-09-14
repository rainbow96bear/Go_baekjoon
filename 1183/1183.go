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

	time := make([]int,N)
	for i:=0 ; i<N ; i++ {
		var A, B int
		fmt.Fscan(reader, &A, &B)
		time[i] = A-B
	}
	sort.Ints(time)
	answer:=0
	if len(time) % 2 == 1{
		answer = 1
	} else {
		answer = time[(len(time)/2)] - time[(len(time)/2)-1] + 1
	}
	fmt.Fprint(writer,answer)
}