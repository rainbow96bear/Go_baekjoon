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

	check := make(map[int]bool)
	answer := 0
	for i:=0 ; i<10 ; i++{
		var input int
		fmt.Fscan(reader, &input)
		input %= 42
		if check[input] == false {
			check[input] = true
			answer++
		}
	}
	fmt.Fprint(writer,answer)
}