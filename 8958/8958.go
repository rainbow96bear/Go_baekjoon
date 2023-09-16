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

	var testCount int
	fmt.Fscan(reader, &testCount)

	for i:=0 ; i<testCount ; i++ {
		var result string
		fmt.Fscan(reader, &result)
		score := 0
		point := 1
		for _, v := range result {
			if string(v) == "O" {
				score += point
				point++
			}else {
				point = 1
			}
		}
		fmt.Fprintln(writer, score)
	}
}