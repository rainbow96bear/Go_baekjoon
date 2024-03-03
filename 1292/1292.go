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

	var A, B int
	fmt.Fscanf(reader, "%d %d\n", &A, &B)

	total := 0
	answer := 0
	for i:=1 ; total < B ; i++ {
		for j:=0 ; j<i ; j++ {
			total++
			if A <= total && total <= B {
				answer += i
			}
		}
	}
	fmt.Fprintln(writer, answer)
}