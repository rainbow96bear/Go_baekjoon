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

	rule := make([]int, N)
	for i:=0 ; i<N ; i++{
		fmt.Fscan(reader, &rule[i])
	}

	line := []int{N}
	for i:=N-1 ; i>0 ; i-- {
		cnt := 0
		for j:=0 ; j<len(line) ; j++{
			if cnt == rule[i-1]{
				break
			} else if line[j] > i {
				cnt++
			}
		}
		line = append(line[:cnt], append([]int{i},line[cnt:]...)...)
	} 
	for _, v := range line {
		fmt.Fprintf(writer, "%d ", v)
	}
}