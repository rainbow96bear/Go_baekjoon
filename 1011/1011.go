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

	var T int
	fmt.Fscanf(reader, "%d\n", &T)

	for i:=0 ; i<T ; i++ {
		var start, end int
		fmt.Fscanf(reader, "%d %d\n", &start, &end)
		result := FindMinMove(0, end-start)
		fmt.Fprintln(writer, result)
	}
}

func FindMinMove(start, end int) int {
	maxLength := 1
	maxMove := 1
	count := 1
	for  ; maxLength < end ; count++ {
		if count%2 == 0 {
			maxMove++
		}
		maxLength += maxMove
	}
	return count
}