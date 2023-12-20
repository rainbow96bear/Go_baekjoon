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

	var N int
	fmt.Fscanf(reader, "%d\n", &N)

	var base []string
	input := ""
	fmt.Fscanf(reader, "%s\n", &input)
	base = strings.Split(input, "")
	for i:=1 ; i<N ; i++ {
		newInput := ""
		fmt.Fscanf(reader, "%s\n", &newInput)
		for i, v := range newInput {
			if base[i] != string(v) {
				base[i] = "?"
			}
		}
	}
	result := strings.Join(base, "")
	fmt.Fprintln(writer, result)
}