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
	seat := make(map[int]bool, N)
	
	answer := 0
	for i:=0 ; i<N ; i++ {
		var num int
		fmt.Fscan(reader, &num)
		if !seat[num] {
			seat[num] = true
		}else {
			answer++
		}
	}
	fmt.Fprintln(writer, answer)
}