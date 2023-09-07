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
	sli := make([]string, N)
	for i:=0 ; i<N ; i++{
		fmt.Fscan(reader,&sli[i])
	}
	sort.Slice(sli,func(i,j int)bool{
		if len(sli[i]) == len(sli[j]){
			return sli[i]<sli[j]
		}
		return len(sli[i]) < len(sli[j])
	})
	preWord := ""
	for _, v := range sli {
		if preWord == v {
			continue
		}
			fmt.Fprintln(writer, v)
			preWord = v
	}
}