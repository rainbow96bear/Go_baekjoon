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

	list := make([]int, 3)
	fmt.Fscanf(reader, "%d %d %d\n", &list[0], &list[1], &list[2])
	sort.Slice(list, func(i, j int)bool{return list[i] < list[j]})

	for _, v := range list {
		fmt.Fprintf(writer, "%d ", v)
	}
}