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

	var N, K int
	fmt.Fscanf(reader, "%d %d\n", &N, &K)

	list := make([]int, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscanf(reader, "%d ", &list[i])
	}
	gap := []int{}
	for i:=1 ; i<N ; i++ {
		gap = append(gap, list[i] - list[i-1])
	}
	sort.Slice(gap, func(i, j int)bool{return gap[i] < gap[j]})
	total := 0
	for _, v := range gap[:len(gap)+1-K] {
		total += v
	}
	fmt.Fprintln(writer, total)
}