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

	answer := Rev(Rev(A) + Rev(B))
	fmt.Fprintln(writer, answer)
}

func Rev(num int) (result int) {
	for num > 0 {
		result += num%10
		num /= 10
		if num > 0 {
			result *= 10
		}
	}
	return result
}