package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var input string
	fmt.Fscan(reader, &input)

	input = strings.ToUpper(input)
	check := make(map[string]int)
	for _, v := range input {
		check[string(v)]++
	}
	var answer string
	var max int
	for i,v := range check{
		if max < v {
			max = v
			answer = i
		} else if max == v {
			answer = "?"
		}
	}
	fmt.Fprint(writer,answer)
}