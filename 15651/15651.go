package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	N, M int
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main(){
	
	defer writer.Flush()

	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	for i:=1; i<=N ; i++ {
		dfs(1, fmt.Sprintf("%d", i))
	}
}

func dfs(deep int, answer string) {
	if deep == M {
		fmt.Fprintln(writer, answer)
		return
	}

	for i:=1 ; i<=N ; i++ {
		dfs(deep+1, fmt.Sprintf("%s %d", answer, i))
	}
} 