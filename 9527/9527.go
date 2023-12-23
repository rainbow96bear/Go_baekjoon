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

	var A, B int64
	fmt.Fscanf(reader, "%d %d", &A, &B)

	result1 := FindOneCount(A-1)
	result2 := FindOneCount(B)
	fmt.Fprintln(writer, result2-result1)
}

func FindOneCount(num int64) int64 {
	if num <= 0 {
		return 0
	}
	MaxPow := int64(0)
	Value := int64(1)
	total :=int64(0)

	for Value*2 <= num {
		MaxPow++
		Value*=2
	}
	total += MaxPow*Value/2
	next := num - Value
	total += next + FindOneCount(next) + 1

	return total
}