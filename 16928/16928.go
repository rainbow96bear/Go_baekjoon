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

	var N, M int
	fmt.Fscan(reader, &N, &M)
	
	skipList := make(map[int]bool)

	ladderList := make(map[int]int)
	for i:=0 ; i<N ; i++ {
		var from, to int
		fmt.Fscan(reader, &from, &to)
		ladderList[from] = to
		skipList[from] = true
	}

	snakeList := make(map[int]int)
	for i:=0 ; i<M ; i++ {
		var from, to int
		fmt.Fscan(reader, &from, &to)
		snakeList[from] = to
		skipList[from] = true
	}


	board := make(map[int]int)
	board[1] = 0
	for i:=2 ; i<=100 ; i++ {
		board[i] = 100
	}
	queue := []int{1}
	for len(queue) > 0 {
		now := queue[0]
		queue = queue[1:]
		if skipList[now] == true {
			continue
		}
		for i:=1 ; i<=6 ; i++ {
			num := board[now]+1
			if ladderList[now+i] > 0 {
				if board[ladderList[now+i]] > num {
					board[ladderList[now+i]] = num
					queue = append(queue, ladderList[now+i])
				}
			}else if snakeList[now+i] > 0 {
				if board[snakeList[now+i]] > num {
					board[snakeList[now+i]] = num
					queue = append(queue, snakeList[now+i])
				}
			} else {
				if board[now+i] > num {
					board[now+i] = num
					queue = append(queue, now+i)
				}
			}
		}
	}
	fmt.Fprintln(writer, board[100])
}