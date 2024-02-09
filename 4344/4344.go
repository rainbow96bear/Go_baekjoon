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

	var C int
	fmt.Fscanf(reader, "%d\n", &C)

	for i:=0 ; i<C ; i++ {
		var N int
		fmt.Fscanf(reader, "%d", &N)

		list := make([]int, N)
		total := 0
		count := 0
		for n:=0 ; n<N ; n++{
			fmt.Fscanf(reader, "%d", &list[n])
			total += list[n]
		}
		for n:=0 ; n<N ; n++{
			if total < list[n]*N {
				count++
			}
		}
		fmt.Fprintf(writer, "%.3f%s\n", float64(count*100)/float64(N),"%")
		fmt.Fscanf(reader, "\n")
	}
}