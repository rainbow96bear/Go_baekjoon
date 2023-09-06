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

	var N, k int
	fmt.Fscan(reader, &N, &k)
	sli := make([]int,N)
	for i:=0 ; i<N ; i++{
		fmt.Fscan(reader, &sli[i])
	}
	sort.Ints(sli)
	fmt.Fprint(writer,sli[len(sli)-k])
}