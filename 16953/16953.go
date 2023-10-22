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

	var a, b int
	fmt.Fscan(reader, &a, &b)

	check := make(map[int]int)
	queue := []int{a}
	for len(queue) > 0 {
		now := queue[0]
		queue = queue[1:]

		next1 := multi(now)
		next2 := addOne(now)
		if check[next1] == 0 {
			check[next1] = check[now]+1
			if next1 <= b {
				queue = append(queue, next1)
			}
		}
		if check[next2] == 0 {
			check[next2] = check[now]+1
			if next2 <= b {
				queue = append(queue, next2)
			}
		}
	}
	if check[b] == 0 {
		fmt.Fprintln(writer, -1)
	}else {
		fmt.Fprintln(writer, check[b]+1)
	}
}
func multi(a int) int {
	return a*2
}
func addOne(a int) int {
	return a*10 + 1
}