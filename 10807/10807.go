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

	var length int
	fmt.Fscan(reader, &length)

	arr := make([]int,length)
	for i:=0 ; i<length ; i++{
		fmt.Fscan(reader, &arr[i])
	}
	var findNum int
	fmt.Fscan(reader, &findNum)
	var answer int
	for i:=0 ; i<length ; i++{
		if findNum == arr[i] {
			answer++
		}
	}
	fmt.Fprint(writer,answer)
}