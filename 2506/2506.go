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
	
	point := 0
	answer := 0
	for i:=0 ; i<N ; i++ {
		var input int
		fmt.Fscanf(reader, "%d", &input)
		if input == 1 {
			point++
			answer += point
		}else {
			point = 0
		}
	}
	fmt.Fprintln(writer, answer)
}