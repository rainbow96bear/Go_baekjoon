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

	list := make([]int, N)
	countList := make([]int, 1000001)
	for i:=0 ; i<N ; i++ {
		var input int
		fmt.Fscanf(reader, "%d\n", &input)
		list[i] = input
		countList[input]++
	}

	for i:=0 ; i<N ; i++ {
		count := -1
		for j:=1 ; j*j<=list[i] ; j++ {
			if list[i] % j == 0 {
				count += countList[j]
				count += countList[list[i]/j]
			}
			if j*j == list[i] {
				count -= countList[list[i]/j]
			}
		}
		fmt.Fprintln(writer, count)
	}
}