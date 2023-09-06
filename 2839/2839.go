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
	fmt.Fscan(reader, &N)
	for i:=N/5 ; i>=0 ; i--{
		for j:=0 ; 5*i+3*j<=N ; j++{
			if 5*i+3*j == N{
				fmt.Fprint(writer,i+j)
				return
			}
		}
	}
	fmt.Fprint(writer,-1)
}