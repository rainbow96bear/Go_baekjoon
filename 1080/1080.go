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
	fmt.Fscanln(reader, &N, &M)
	Board := make([][]string,N)

	for i:=0 ; i<N ; i++{
		var input string
		fmt.Fscan(reader, &input)
		Board[i] = strings.Split(input,"")
	}
	for i:=0 ; i<N ; i++{
		var input string
		fmt.Fscan(reader, &input)
		input = strings.TrimSpace(input)
		for j, v := range strings.Split(input,""){
			if Board[i][j] == v {
				Board[i][j] = "0"
			} else {
				Board[i][j] = "1"
			}
		}
	}

	cnt:=0
	for i:=0 ; i<N-2 ; i++{
		for j:=0 ; j<M-2 ; j++{
			if Board[i][j] == "1" {
				reverse(Board, i, j)
				cnt++
			}
		}
	}

	for i:=0 ; i<N ; i++{
		for j:=0 ; j<M ; j++{
			if Board[i][j] == "1" {
				fmt.Fprintln(writer, -1)
				return
			}
		}
	}
	fmt.Fprintln(writer, cnt)
}

func reverse(Board [][]string, idxR, idxC int){
	for i:=idxR ; i<idxR+3 ; i++{
		for j:=idxC ; j<idxC+3 ; j++{
			if Board[i][j] == "0" {
				Board[i][j] = "1"
			}else {
				Board[i][j] = "0"
			}
		}
	}
}