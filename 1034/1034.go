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

	var N, M int
	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	rowNums := make(map[string]int)
	rowZero := make(map[string]int)
	for i:=0 ; i<N ; i++ {
		var input string
		fmt.Fscanf(reader, "%s\n", &input)
		if rowNums[input] == 0 {
			count := 0
			for _, v := range input {
				if v == '0' {
					count++
				}
			}
			rowZero[input] = count
		}
		rowNums[input]++
	}

	var K int
	fmt.Fscanf(reader, "%d\n", &K)
	max := 0
	for i, v:= range rowZero {
		if K >= v && (K-v)%2 == 0 {
			if max < rowNums[i] {
				max = rowNums[i]
			}
		}
	}
	fmt.Fprintln(writer, max)
}