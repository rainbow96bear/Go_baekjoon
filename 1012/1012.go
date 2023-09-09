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

	var T int
	fmt.Fscan(reader, &T)
	for i:=0 ; i<T ; i++ {
		var N, M, K int
		fmt.Fscan(reader, &N, &M, &K)
		board := make([][]int, N)
		
		for j:=0 ; j<N ; j++ {
			board[j] = make([]int,M)
		}

		for j:=0 ; j<K ; j++ {
			var x, y int
			fmt.Fscan(reader, &x, &y)
			board[x][y] = 1
		}

		cnt := 0
		for j:=0 ; j<N ; j++{
			for k:=0 ; k<M ; k++{
				if board[j][k] == 1 {
					cnt++
					DFS(board, j, k, N, M)
				}
			}
		}
		fmt.Fprintln(writer,cnt)
	}
}

func DFS(board [][]int, x, y, N, M int){
	stack := [][]int{{x,y}}
	for len(stack)>0 {
		var X, Y int = stack[len(stack)-1][0], stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]

		if X-1 >= 0 && board[X-1][Y] == 1 {
			board[X-1][Y]=0
			stack = append(stack, []int{X-1,Y})
		}
		if X+1 < N && board[X+1][Y] == 1 {
			board[X+1][Y]=0
			stack = append(stack, []int{X+1,Y})
		}
		if Y-1 >= 0 && board[X][Y-1] == 1 {
			board[X][Y-1]=0
			stack = append(stack, []int{X,Y-1})
		}
		if Y+1 < M && board[X][Y+1] == 1 {
			board[X][Y+1]=0
			stack = append(stack, []int{X,Y+1})
		}
	}
}