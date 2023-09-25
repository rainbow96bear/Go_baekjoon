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

	cntBlue := 0
	cntWhite := 0
	check := make([][]bool,N)
	
	board := make([][]int, N)
	for i:=0 ; i<N ; i++ {
		for j:=0 ; j<N ; j++{
			var input int
			fmt.Fscan(reader, &input)
			board[i] = append(board[i], input)
			check[i] = append(check[i], false)
		}
	}
	length := N
	for length > 0 {

		for X:=0 ; X <= N-length ; X += length {
			for Y:=0 ; Y <= N-length ; Y += length {
				if check[X][Y] == true {
					continue
				}

				color := 0
				if board[X][Y] == 1 {
					color = 1
				}
				checkSame := true
				
				for i:=0 ; i<length ; i++ {
					for j:=0 ; j<length ; j++ {
						check[X+i][Y+j] = true
						if board[X+i][Y+j] != color{
							checkSame = false
							break
						}
					}
					if checkSame == false {
						break
					}
				}
				
				if checkSame == false {
					for i:=0 ; i<length ; i++ {
						for j:=0 ; j<length ; j++ {
							check[X+i][Y+j] = false
						}
					}
				}else {
					if color == 0 {
						cntWhite++
					}else {
						cntBlue++
					}
				}
			} 
		}
		length /= 2

	}
	fmt.Fprintf(writer, "%d\n%d", cntWhite, cntBlue)
}