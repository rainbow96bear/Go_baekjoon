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
	fmt.Fscanf(reader, "%d\n", &N)

	arr := make([]int, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscanf(reader, "%d\n", &arr[i])
	}

	sort.Slice(arr, func(i, j int)(bool){return arr[i] < arr[j]})
	answer := 5
	for i:=0 ; i<N ; i++ {
		needNum := 4
		for j:= i+1 ; j<i+5&&j<N ; j++ {
			if(arr[j] - arr[i]) < 5 {
				needNum--
			}
		}
		if needNum < answer {
			answer = needNum
		}
	}
	fmt.Fprintln(writer, answer)
}