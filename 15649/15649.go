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

	isVisit := make([]bool, N+1)

	for i:=1 ; i<=N ; i++ {
		isVisit[i] = true
		dfs(isVisit, 1, fmt.Sprintf("%d", i))
		isVisit[i] = false
	}
}

func dfs(isVisit []bool, deep int, answer string) {
	if M == deep {
		fmt.Fprintln(writer, answer)
		return
	}
	for i:=1 ; i<=N ; i++ {
		if isVisit[i] == false {
			isVisit[i] = true
			dfs(isVisit, deep+1, fmt.Sprintf("%s %d", answer, i))
			isVisit[i] = false
		}
	}
}