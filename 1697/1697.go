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
	if N == K {
		fmt.Fprintln(writer, 0)
		return
	}

	queue := []int{N}
	check := make(map[int]int)
	cnt := 1
	for {
		newQueue := []int{}
		for i:= 0 ; i<len(queue) ; i++ {
			nowNum := queue[i]
			if nowNum-1 == K || nowNum+1 == K || nowNum*2 == K {
				fmt.Fprintln(writer, cnt)
				return
			}
			if check[nowNum-1] == 0 && nowNum-1 >= 0 {
				check[nowNum-1] = cnt
				newQueue = append(newQueue, nowNum-1)
			}
			if check[nowNum+1] == 0 && nowNum-1 <= 100000 {
				check[nowNum+1] = cnt
				newQueue = append(newQueue, nowNum+1)
			}
			if check[nowNum*2] == 0 && nowNum*2 <= 100000 {
				check[nowNum*2] = cnt
				newQueue = append(newQueue, nowNum*2)
			}
		}
		queue = newQueue
		cnt++
	}
}