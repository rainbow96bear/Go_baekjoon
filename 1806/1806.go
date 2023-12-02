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

	var N, S int
	fmt.Fscanf(reader, "%d %d\n",&N, &S)

	list := make([]int, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscanf(reader, "%d", &list[i])
	}
	result := Process(list, S)
	fmt.Fprintln(writer, result)
}

func Process(list []int, S int) int {
	if list[0] >= S {
		return 1
	}
	start, end := 0, 0
	length := 1<<30
	total := list[start]
	for start <= end && end < len(list){
		if total < S {
			end++
			if end < len(list) {
				total += list[end]
			}else {
				break
			}
		}else if total >= S {
			if length > end-start+1 {
				length = end-start+1
			}
			total -= list[start]
			start++
		}
	}
	if length == 1<<30 {
		return 0
	}
	return length
}