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

	var N, M int
	fmt.Fscan(reader, &M, &N)

	for i:=M ; i<=N ; i++{
		if i == 1 {
			continue
		}
		check := false
		for j:=2 ; j*j <= i ; j++{
			if i%j == 0 {
				check = true
				break
			}
		}
		if check == false {
			fmt.Fprintln(writer, i)
		}
	}
}