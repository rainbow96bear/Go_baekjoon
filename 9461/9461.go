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

	P := []int{1,1,1,2,2}

	var N int
	fmt.Fscan(reader, &N)

	max := 0
	requestList := []int{}
	for i:=0 ; i<N ; i++ {
		var input int
		fmt.Fscan(reader, &input)
		requestList = append(requestList, input)
		if max < input {
			max = input
		}
	}
	for i:=5 ; i<max ; i++ {
		P = append(P, P[i-1]+P[i-5])
	}
	for _, v := range requestList {
		fmt.Fprintln(writer,P[v-1])
	}
}