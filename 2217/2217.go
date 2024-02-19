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
	fmt.Fscanf(reader, "%d\n", &N)

	list := make([]int, N)

	for i:=0 ; i<N ; i++ {
		fmt.Fscanf(reader, "%d\n", &list[i])
	}
	sort.Slice(list, func(i, j int)bool{return list[i] > list[j]})
	max := 0
	for i, v := range list {
		if max < v*(i+1) {
			max = v*(i+1)
		}
	}
	fmt.Fprintln(writer, max)
}