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
	fmt.Fscanf(reader, "%d\n", &T)

	for i:=0 ; i<T ; i++ {
		var d int
		fmt.Fscanf(reader, "%d\n", &d)
		t := 0
		for {
			isPossible := t+1
			if isPossible+isPossible*isPossible <= d {
				t++
			}else {
				break
			}
		}
		fmt.Fprintln(writer, t)
	}
}