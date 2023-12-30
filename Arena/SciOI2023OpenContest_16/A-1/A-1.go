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

	var T int
	fmt.Fscanf(reader, "%d\n", &T)

	for t:=0 ; t<T ; t++ {
		var N, M, K int
		fmt.Fscanf(reader, "%d %d %d\n", &N, &M, &K)
		result := FindMinCount(N, M, K)
		fmt.Fprintln(writer, result)
	}
}

func FindMinCount(N, M, K int) (count int) {
	count = 0
	if N <= M*K {
		return 1
	}
	if M*K < 2 {
		return -1
	}
	for {
		N -= M*K
		count++
		if N <= 0 {
			break
		}
		N++
		count++
	}
	return count
}