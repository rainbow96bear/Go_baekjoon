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

	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	result := findMinCoin(n)
	if n == result {
		fmt.Fprintln(writer, "-1")
	}else {
		fmt.Fprintln(writer, result)
	}
}

func findMinCoin(n int) int {
	min := n
	for i:=0 ; 5*i<=n ; i++{
		if (n-(5*i))%2 == 0 && min > i+((n-(5*i))/2){
			min = i+((n-(5*i))/2)
		}
	}
	return min
}