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

	var N int
	fmt.Fscanf(reader, "%d\n", &N)

	check := make(map[int64]int)
	for i:=0 ; i<N ; i++ {
		var input int64
		fmt.Fscanf(reader, "%d\n", &input)
		check[input]++
	}
	var answer int64
	count := 0

	for i, v := range check {
		if count < v {
			count = v
			answer = i
		} else if count == v {
			if answer > i {
				answer = i
			}
		}
	}

	fmt.Fprintln(writer, answer)
}