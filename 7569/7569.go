package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	x int
	y int
	z int
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var M, N, H int
	fmt.Fscan(reader, &M, &N, &H)

	box := make([][][]int,H)
	queue := []position{}
	cnt := 0
	for i:=0 ; i<H ; i++ {
		box[i] = make([][]int, N)
		for j:=0 ; j<N ; j++ {
			box[i][j] = make([]int, M)
			for k:=0; k<M ; k++ {
				fmt.Fscan(reader, &box[i][j][k])
				if box[i][j][k] == 1 {
					queue = append(queue, position{x:k, y:j, z:i})
				} else if box[i][j][k] == 0 {
					cnt++
				}
			}
		}
	}
	if cnt == 0 {
		fmt.Fprintln(writer, 0)
		return
	}

	answer := 0
	for len(queue)>0 {
		newQueue := []position{}
		for len(queue)>0 {
			now := queue[0]
			queue = queue[1:]
			dx := []int{1, -1, 0, 0, 0, 0}
			dy := []int{0, 0, 1, -1, 0, 0}
			dz := []int{0, 0, 0, 0, 1, -1}
	
			for i:=0 ; i<6 ; i++ {
				X := now.x + dx[i]
				Y := now.y + dy[i]
				Z := now.z + dz[i]
	
				if 0 <= X && X < M && 0 <= Y && Y < N && 0 <= Z && Z < H {
					if box[Z][Y][X] == 0 {
						newQueue = append(newQueue, position{x:X, y:Y, z:Z})
						box[Z][Y][X] = 1
						cnt--
					}
				}
			}
		}
		queue = newQueue
		if len(queue) > 0 {
			answer++
		}
	}

	if cnt > 0 {
		fmt.Fprintln(writer, -1)
	}else {
		fmt.Fprintln(writer, answer)
	}
}