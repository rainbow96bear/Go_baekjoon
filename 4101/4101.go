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

	for {
		var a, b int
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		if a==0 && b==0 {
			break
		}
		if a > b {
			fmt.Fprintln(writer, "Yes")
		}else {
			fmt.Fprintln(writer,"No")
		}
	}
}