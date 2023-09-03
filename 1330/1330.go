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

	var a, b int
	fmt.Fscan(reader, &a, &b)
	switch {
	case a>b :
		fmt.Fprintln(writer,">")
	case a<b :
		fmt.Fprintln(writer,"<")
	case a==b :
		fmt.Fprintln(writer,"==")
	}
}