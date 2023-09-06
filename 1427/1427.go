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

	var N string
	fmt.Fscan(reader, &N)
	sli := make([]int,len(N))
	for i,v := range N {
		sli[i] = int(byte(v) -'0')
	}
	sort.Ints(sli)
	for i,_ := range sli {
		fmt.Fprint(writer,sli[len(sli)-1-i])
	}
}