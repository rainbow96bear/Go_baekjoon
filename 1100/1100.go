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
	total := 0
	for i:=0 ; i<8 ; i++ {
		var input string
		fmt.Fscanf(reader, "%s\n", &input)
		list := strings.Split(input, "")
		for j, value := range list {
			if (i+j) % 2 == 0 && value == "F"{
				total ++
			}
		}
	}
	fmt.Fprintln(writer, total)
}