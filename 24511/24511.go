package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)

	checkQS := make([]int,N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscan(reader, &checkQS[i])
	}

	newDeque := list.New()
	for i:=0 ; i<N ; i++{
		var input int
		fmt.Fscan(reader, &input)
		if checkQS[i] == 0 {
			newDeque.PushFront(input)
		}
	}
	var M int
	fmt.Fscan(reader, &M)

	for i:=0 ; i<M ; i++{
		var input int
		fmt.Fscan(reader, &input)
		newDeque.PushBack(input)
		fmt.Fprintf(writer,"%d ", newDeque.Front().Value)
		newDeque.Remove(newDeque.Front())
	}
}