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

	var N, m, M, T, R int
	fmt.Fscanf(reader, "%d %d %d %d %d\n", &N, &m, &M, &T, &R)

	total := 0
	bloodBounce := m
	if m + T <= M {
		for execTime := 0 ; execTime < N ; total++{
			if bloodBounce + T <= M {
				execTime ++
				bloodBounce += T
			}else {
				bloodBounce -= R
				if bloodBounce < m{
					bloodBounce = m
				}
			}
		}
	}else {
		total = -1
	}

	fmt.Fprintln(writer, total)
}