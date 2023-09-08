package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N, K int
	fmt.Fscan(reader, &N, &K)

	list := make([]int,N)
	for i:=1 ; i<=N ; i++{
		list[i-1]=i
	}
	fmt.Fprintf(writer,"<")
	for len(list)>1 {
		if K%len(list) == 0 {
			fmt.Fprintf(writer, "%d, ", list[len(list)-1])
			list = list[:len(list)-1]
		}else {
			fmt.Fprintf(writer, "%d, ", list[K%len(list)-1])
			list = append(list[K%len(list):],list[:K%len(list)-1]...)
		}
	}
	fmt.Fprintf(writer,"%d>",list[0])
}