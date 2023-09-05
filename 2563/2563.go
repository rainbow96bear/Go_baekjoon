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

	var papersNum int
	fmt.Fscan(reader, &papersNum)

	board := [100][100]int{}
	answer := 0
	for i:=0 ; i<papersNum ; i++{
		var x, y int
		fmt.Fscan(reader, &x, &y)
		for j:=x ; j<x+10 ; j++ {
			for k:=y ; k<y+10 ; k++ {
				if board[j][k] == 0{
					board[j][k] = 1
					answer++
				}
			}
		}
	}
	fmt.Fprint(writer,answer)
}