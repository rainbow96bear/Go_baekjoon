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

	var L int
	fmt.Fscan(reader, &L)

	groupS := make([]int, L)
	for i:=0 ; i<L ; i++{
		fmt.Fscan(reader, &groupS[i])
	}
	sort.Ints(groupS)
	var n int
	fmt.Fscan(reader, &n)
	if n < groupS[0]{
		fmt.Fprint(writer, (n * (groupS[0] - n))-1)
		return
	}
	for i:=1 ; i<L ; i++{
		if groupS[i-1] < n && n < groupS[i] {
			fmt.Fprint(writer, (n - groupS[i-1]) * (groupS[i] - n) -1)
			return
		}
	}
	fmt.Fprint(writer, 0)
}