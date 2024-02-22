package main

import (
	"bufio"
	"fmt"
	"os"
)
type Egg struct {
	hard int
	weight int
}

var (
	answer = 0
)
func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscanf(reader, "%d\n", &N)

	eggs := make([]Egg, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscanf(reader, "%d %d\n", &eggs[i].hard, &eggs[i].weight)
	}
	dfs(0, N, 0, eggs)
	fmt.Fprintln(writer, answer)
}

func dfs(num, N, count int, eggs []Egg){
	if answer < count {
		answer = count
	}
	if num == N {
		return
	}
	if eggs[num].hard <= 0 {
		dfs(num+1, N, count, eggs)
	}else {

		for i:=0 ; i<N ; i++ {
			if eggs[i].hard <= 0 {
				continue
			}
			if i != num && eggs[i].hard > 0 {
				add := 0
				eggs[num].hard -= eggs[i].weight
				eggs[i].hard -= eggs[num].weight
				if eggs[num].hard <= 0 {
					add++
				}
				if eggs[i].hard <= 0 {
					add++
				}
				
				dfs(num+1, N, count + add, eggs)
				eggs[num].hard += eggs[i].weight
				eggs[i].hard += eggs[num].weight
			}
		}
	}

}