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

	for t:=0 ; t<T ; t++ {
		var N int
		fmt.Fscanf(reader, "%d\n", &N)
		
		list := make(map[int]bool)

		for i:=0 ; i<N ; i++ {
			var num int
			fmt.Fscanf(reader, "%d", &num)
			list[num] = true
		}
		fmt.Fscanf(reader, "\n")
		var M int
		fmt.Fscanf(reader, "%d\n", &M)

		for i:=0 ; i<M ; i++ {
			var num int
			fmt.Fscanf(reader, "%d", &num)
			if list[num] {
				fmt.Fprintln(writer, 1)
			}else {
				fmt.Fprintln(writer, 0)
			}
		}
		fmt.Fscanf(reader, "\n")
	}
}