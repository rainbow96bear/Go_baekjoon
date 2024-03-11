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

	odd := 0
	even := 0
	for i:=0 ; i<N ; i++ {
		var input int
		fmt.Fscanf(reader, "%d", &input)

		even += input/2
		if input % 2 == 1 {
			odd++
		}
	}
	if (even-odd) % 3 != 0 || even < odd {
		fmt.Fprintln(writer, "NO")
	}else {
		fmt.Fprintln(writer, "YES")
	}
}