package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N, M int
	fmt.Fscanln(reader, &N ,&M)
	maxLength:=N
	if N > M {
		maxLength = M
	}
	board := make([][]string, N)
	for i:=0 ; i<N ; i++{
		input, _ := reader.ReadString('\n')
		strings.TrimSpace(input)
		for _, v := range input{
			board[i] = append(board[i], string(v))
		}
	}
	
	for i:=maxLength-1 ; i>0 ; i-- {
		for j:=0 ; j<N-i ; j++{
			for k:=0 ; k<M-i ; k++ {
				if board[j][k] == board[j+i][k] && board[j][k] == board[j][k+i] && board[j][k] == board[j+i][k+i] {
					fmt.Fprint(writer,(i+1)*(i+1))
					return
				}
			}
		}
	}
	fmt.Fprint(writer,1)
}