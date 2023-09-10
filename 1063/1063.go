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

	var (
		King, Rock string
		N int
	)

	K := make(map[string]int)
	R := make(map[string]int)
	fmt.Fscan(reader, &King, &Rock, &N)
	K["x"] = int(King[0]-'A')
	K["y"] = int(King[1]-'1')
	R["x"] = int(Rock[0]-'A')
	R["y"] = int(Rock[1]-'1')

	for i:=0 ; i<N ; i++{
		var command string
		fmt.Fscan(reader, &command)
		switch command {
		case "R" :
			move(K, R, 1, 0)
		case "L" :
			move(K, R, -1, 0)
		case "B" :
			move(K, R, 0, -1)
		case "T" :
			move(K, R, 0, 1)
		case "RT" :
			move(K, R, 1, 1)
		case "LT" :
			move(K, R, -1, 1)
		case "RB" :
			move(K, R, 1, -1)
		case "LB" :
			move(K, R, -1, -1)
		}
	}
	fmt.Fprintf(writer, "%s%s\n",string(byte(int('A')+K["x"])), string(byte(int('1')+K["y"])))
	fmt.Fprintf(writer, "%s%s\n",string(byte(int('A')+R["x"])), string(byte(int('1')+R["y"])))
}
func move(K, R map[string]int, mx, my int){
	x, y := "x", "y"
	if 0 <= K[x]+mx && K[x]+mx <= 7 && 0 <= K[y]+my && K[y]+my <= 7{
		if K[x]+mx == R[x] && K[y]+my == R[y]{
			if 0 <= R[x]+mx && R[x]+mx <= 7 && 0 <= R[y]+my && R[y]+my <= 7 {
				K[x]+=mx
				K[y]+=my
				R[x]+=mx
				R[y]+=my
			}
		}else {
			K[x]+=mx
			K[y]+=my
		}
	}
}