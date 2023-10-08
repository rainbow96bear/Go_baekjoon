package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct{
	x int
	y int
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var M, N int
	fmt.Fscan(reader, &M, &N)

	box := make([][]int, N)

	cnt := 0
	queue := []position{}
	answer := 0
	for i:=0 ; i<N ; i++ {
		box[i] = make([]int, M)
		for j:=0 ; j<M ; j++ {
			fmt.Fscan(reader, &box[i][j])
			if box[i][j] == 1 {
				queue = append(queue, position{x:i, y:j})
			} else if box[i][j] == 0 {
				cnt++
			}
		}
	}
	if cnt == 0 {
		fmt.Fprintln(writer, 0)
		return
	}
	for len(queue)>0 {
		newQueue := []position{}
		for len(queue)>0 {
			now := queue[0]
			queue = queue[1:]
			
			dx := []int{0, 0, 1, -1}
			dy := []int{-1, 1, 0, 0}

			for i:=0 ; i<4 ; i++ {
				X := now.x + dx[i]
				Y := now.y + dy[i]

				if 0 <= X && X < N && 0 <= Y && Y < M {
					if box[X][Y] == 0 {
						newQueue = append(newQueue, position{x:X, y:Y})
						cnt--
						box[X][Y] = 1
					}
				}
			}

		}
		queue = newQueue
		if len(queue)>0 {
			answer++
		}
	}

	if cnt > 0 {
		fmt.Fprintln(writer, -1)
	}else {
		fmt.Fprintln(writer, answer)
	}
}