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

	var H, W int
	fmt.Fscanf(reader, "%d %d\n", &H, &W)

	blockHeight := make([]int, W)

	for i, _ := range blockHeight {
		fmt.Fscanf(reader, "%d", &blockHeight[i])
	}
	fmt.Fscanf(reader, "\n")


	total := 0
	for i:=1 ; i<W-1 ; i++ {
		leftmax, rightmax := blockHeight[i], blockHeight[i]
		for index:=i ; index>=0 ; index-- {
			if leftmax < blockHeight[index] {
				leftmax = blockHeight[index]
			}
		}
		for index:=i ; index<len(blockHeight) ; index++ {
			if rightmax < blockHeight[index] {
				rightmax = blockHeight[index]
			}
		}
		
		minHeight := min(leftmax, rightmax)
		if minHeight-blockHeight[i] > 0 {
			total += minHeight-blockHeight[i]
		}
	}
	fmt.Fprintln(writer, total)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}