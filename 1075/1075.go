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

	var N, F int
	fmt.Fscanf(reader, "%d\n", &N)
	fmt.Fscanf(reader, "%d\n", &F)
	N = (N/100) * 100
	for i:=0 ; i<100; i++ {
		if (N+i) % F == 0 {
			if i<10 {
				answer := fmt.Sprintf("0%d",i)
				fmt.Fprintln(writer, answer)
				return
			}
			fmt.Fprintln(writer, i)
			return
		}
	}
}