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

	var N, K int
	fmt.Fscan(reader, &N, &K)
	cnt := 0
	for i:=1 ; i<=N ; i++{
		if N%i == 0{
			cnt++
			if cnt == K {
				fmt.Fprint(writer,i)
				return
			}
		}
	}
	fmt.Fprint(writer,0)
}