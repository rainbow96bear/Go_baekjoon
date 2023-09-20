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

	var X int
	fmt.Fscan(reader, &X)
	
	check := make(map[int]int)
	check[1] = 0
	for i:=1 ; i<=X ; i++ {
		if check[i+1] == 0 || check[i+1] > check[i]+1{
			check[i+1] = check[i]+1
		}
		if check[i*2] == 0 || check[i*2] > check[i]+1{
			check[i*2] = check[i]+1
		}
		if check[i*3] == 0 || check[i*3] > check[i]+1{
			check[i*3] = check[i]+1
		}
	}
	fmt.Fprintln(writer, check[X])
}