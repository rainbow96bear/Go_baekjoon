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

	input := make([]string,5)
	max := 0
	for i:=0 ; i<5 ; i++{
		fmt.Fscan(reader, &input[i])
		if max < len(input[i]){
			max = len(input[i])
		}
	}
	for i:=0 ; i<max ; i++{
		for j:=0 ; j<5 ; j++{
			if len(input[j])>i{
				fmt.Fprint(writer,string(input[j][i]))
			}
		}
	}
}