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
	sli1 := make([]int, N)
	sli2 := []int{}
	for i:=0 ; i<N ; i++{
		fmt.Fscan(reader, &sli1[i])
	}
	sli2 = append(sli2,sli1...)
	sort.Ints(sli2)
    check := make(map[int]int)
	cnt := 0
	for i, v := range sli2 {
		if i == 0{
			check[v] = i-cnt
		}else {
			if sli2[i-1] == v{
				cnt++
				continue
			}
			check[v] = i-cnt
		}
	}
	for _,v := range sli1{
		fmt.Fprintf(writer,"%d ",check[v])
	}
}