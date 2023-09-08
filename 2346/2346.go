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

	ballons := make([]int,N)
	checkPop := make([]bool,N)
	for i:=0 ; i<N ; i++{
		fmt.Fscan(reader, &ballons[i])
	}
	index := 0
	for i:=0 ; i<N ; i++ {
		fmt.Fprintf(writer, "%d ", index+1)
		move := ballons[index]
		checkPop[index] = true
		if i == N-1 {
			break
		}
		if move < 0 {
			move = -move
			for j:=0 ; j<move ; j++{
				for {
					index = (index-1+N)%N
					if checkPop[index] == false {
						break
					}
				}
			}
		}else{
			for j:=0 ; j<move ; j++{
				for {
					index = (index+1) % N
					if checkPop[index] == false {
						break
					}
				}
			}
		}
		
	}
}