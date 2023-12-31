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
	result := Factorial(N)
	fmt.Fprintln(writer, result)
}

func Factorial(N int) (result int) {
	result = 1
	for i:=1 ; i<=N ; i++ {
		result *= i
	}
	return result
}