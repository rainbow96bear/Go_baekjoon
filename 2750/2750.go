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
	sli := make([]int,N) 
	for i:=0 ; i<N ; i++{
		fmt.Fscan(reader,&sli[i])
	}
	sort.Ints(sli)
	for i:=0 ; i<N ; i++{
		fmt.Fprintln(writer,sli[i])
	}
}