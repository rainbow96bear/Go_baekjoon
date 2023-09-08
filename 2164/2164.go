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

	var N int
	fmt.Fscan(reader, &N)

	queue := []int{}
	for i:=1 ; i<=N ; i++{
		queue = append(queue, i)
	}
	
	for len(queue)>1{
		queue = queue[1:]
		if len(queue)==1 {
			fmt.Fprint(writer, queue[0])
			return
		}
		queue = append(queue[1:], queue[0])
	}
	fmt.Fprint(writer, queue[0])
}