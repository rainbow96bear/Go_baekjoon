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

	arr := make([]int,N)
	for i:=0 ; i<N ; i++{
		fmt.Fscan(reader,&arr[i])
	}
	sort.Ints(arr)
	fmt.Fprint(writer, arr[0], arr[N-1])
}