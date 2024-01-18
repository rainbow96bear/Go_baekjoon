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

	checkCount := make(map[string]int)
	for i:=0 ; i<N ; i++ {
		var book string
		fmt.Fscanf(reader, "%s\n", &book)
		checkCount[book]++
	}

	result := bestSellor(checkCount)
	fmt.Fprintln(writer, result)
}

func bestSellor(checkCount map[string]int) string {
	maxCount := 0
	book := "" 
	for i, v := range checkCount {
		if v > maxCount {
			book = i
			maxCount = v
		} else if v == maxCount {
			if book > i {
				book = i
			}
		}
	}
	return book
}