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
	fmt.Fscanf(reader, "%d\n", &N)

	list := make([]int, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscanf(reader, "%d ", &list[i])
	}
	max := 0
	for i:=0 ; i<N ; i++ {
		count := 0
		leftLean := 1000000001.0
		rightLean := -1000000001.0
		for j:=i-1 ; j>=0 ; j-- {
			lean := float64(list[i]-list[j])/float64(i-j)
			if leftLean > lean {
				count++
				leftLean = lean
			}
		}
		for j:=i+1 ; j<N ; j++ {
			lean := float64(list[i]-list[j])/float64(i-j)
			if rightLean < lean {
				count++
				rightLean = lean
			}
		}
		
		if max < count {
			max = count
		}
	}
	fmt.Fprintln(writer, max)
}