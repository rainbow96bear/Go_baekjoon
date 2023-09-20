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

	list := make([]int, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscan(reader, &list[i])
	}

	sort.Ints(list)
	answer := 0
	for i:=0 ; i<N ; i++ {
		answer += list[i]*(N-i)
	}
	fmt.Fprintln(writer, answer)
}