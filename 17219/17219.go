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
	fmt.Fscan(reader, &N, &M)

	list := make(map[string]string, N)
	for i:=0 ; i<N ; i++ {
		var address, password string
		fmt.Fscan(reader, &address, &password)
		list[address] = password
	}
	for i:=0 ; i<M ; i++ {
		var address string
		fmt.Fscan(reader, &address)
		fmt.Fprintln(writer, list[address])
	}
}