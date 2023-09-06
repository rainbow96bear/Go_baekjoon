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

	var X, Y int
	fmt.Fscanln(reader, &X, &Y)

	board := make([][]string,X)
	for i:=0 ; i<X ; i++{
		var input string
		input,_ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		board[i] = append(board[i],strings.Split(input, "")...)
	}
	min := 64
	for i:=0 ; i<=X-8 ; i++{
		for j:=0; j<=Y-8 ; j++{
			newBoard := make([][]string,8)
			for k:=0 ; k<8 ; k++{
				newBoard[k] = board[(i+k)][j:j+8]
			}
			rst1 := check("W",newBoard)
			rst2 := check("B",newBoard)
			if rst1 < min || rst2 < min {
				if rst1 < rst2 {
					min = rst1
				}else {
					min = rst2
				}
			}
		}
	}
	fmt.Fprint(writer,min)
}
func check(color string, board [][]string) int {
	min := 64
	cnt := 0
	for i:=0 ; i<8 ; i++{
		for j:=0 ; j<8 ; j++{
			if (i+j) % 2 == 0 {
				if color != board[i][j]{
					cnt++
				}
			}else if (i+j) % 2 == 1 {
				if color == board[i][j]{
					cnt++
				}
			}
		}
	}
	if min > cnt{
		min = cnt
	}
	return min
}