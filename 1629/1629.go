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

	var a, b, c int64
	fmt.Fscan(reader, &a, &b, &c)
	newA := a%c
	answer := int64(1)
	for b>0 {
		if b%2 == 1 {
			answer = (answer*newA) % c
		}
		b /= 2
		newA = (newA*newA) % c
	}
	fmt.Fprintln(writer, answer)
}