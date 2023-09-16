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

	var A, B, C int
	fmt.Fscan(reader, &A, &B, &C)

	multipleResult := A*B*C

	answerList := make([]int,10)

	for multipleResult > 0 {
		Number := multipleResult%10
		answerList[Number]++
		multipleResult /= 10
	}
	for _, v := range answerList {
		fmt.Fprintln(writer, v)
	}
}