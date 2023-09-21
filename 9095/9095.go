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

	var TestCase int
	fmt.Fscan(reader, &TestCase)

	check := make(map[int]int)
	check[1]=1
	check[2]=2
	check[3]=4
	
	inputList := make([]int,TestCase)
	max := 0
	for i:=0 ; i<TestCase ; i++ {
		var n int
		fmt.Fscan(reader, &n)
		if max < n {
			max = n
		}
		inputList[i]=n
	}
	for i:=4 ; i<=max ; i++ {
		check[i]= check[i-1] + check[i-2] + check[i-3]
	}
	for i:=0 ; i<len(inputList) ; i++ {
		fmt.Fprintln(writer, check[inputList[i]])
	}
}