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

	coinPriceList := make([]int, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscan(reader, &coinPriceList[i])
	}

	answer := 0
	for i:=1 ; i<=N ; i ++ {
		answer += K/coinPriceList[N-i]
		K %= coinPriceList[N-i]
	}
	fmt.Fprintln(writer, answer)
}