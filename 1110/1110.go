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

	result := FindCycleLength(N)
	fmt.Fprintln(writer, result)
}

func FindCycleLength(N int) int {
	count := 0
	newNum := N
	lastNum := 0
	for {
		eachSum := AddEachNum(newNum, lastNum)
		newNum = newNum%10*10 + eachSum%10
		count++
		if N == newNum {
			break
		}
	}
	return count
}

func AddEachNum(N, sum int) int {
	result := 0
	N += sum
	for N > 0 {
		result += N%10
		N /= 10
	}
	return result
}