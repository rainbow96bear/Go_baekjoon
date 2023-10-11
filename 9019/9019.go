package main

import (
	"bufio"
	"fmt"
	"os"
)
type Node struct {
	value int
	way string
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var T int
	fmt.Fscan(reader, &T)

	for i:=0 ; i<T ; i++{
		var from, to int
		fmt.Fscan(reader, &from, &to)

		decimals := 1000
		
		result := BFS(from, to, decimals)

		fmt.Fprintln(writer, result)
	}
}

func D(from int) int{
	return (from*2)%10000
}

func S(from int) int{
	if from == 0 {
		return 9999
	}
	return from-1
}

func L(from, decimals int) int{
	return (from % decimals) * 10 + (from / decimals)
}

func R(from, decimals int) int{
	return ((from % 10) * decimals) + (from / 10)
}

func BFS(from, to, decimals int) string{
	command := make(map[int]string)
	queue := []int{from}

	for len(queue) > 0 {
		now := queue[0]
		queue = queue[1:]
		if now == to {
			return command[now]
		}
		afterD := D(now)
		afterS := S(now)
		afterL := L(now, decimals)
		afterR := R(now, decimals)

		if len(command[afterD])>0 == false {
			if afterD == to {
				return command[now]+"D"
			}
			queue = append(queue, afterD)
			command[afterD] = command[now]+"D"
		}
		if len(command[afterS])>0 == false {
			if afterS == to {
				return command[now]+"S"
			}
			queue = append(queue, afterS)
			command[afterS] = command[now]+"S"
		}
		if len(command[afterL])>0 == false {
			if afterL == to {
				return command[now]+"L"
			}
			queue = append(queue, afterL)
			command[afterL] = command[now]+"L"
		}
		if len(command[afterR])>0 == false {
			if afterR == to {
				return command[now]+"R"
			}
			queue = append(queue, afterR)
			command[afterR] = command[now]+"R"
		}
		if len(command[from]) == 0 {
			command[from] = "x"
		}
	}
	return command[to]
}