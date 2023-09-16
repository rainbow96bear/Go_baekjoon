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
	fmt.Fscan(reader, &T)

	for i:=0 ; i<T ; i++ {
		var H, W, N int
		fmt.Fscan(reader, &H, &W, &N)
		fmt.Fprintln(writer, (((N-1)%H)+1)*100+(((N-1)/H)+1))
	}
}