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

	var a, b, c, d, e int
	fmt.Fscanf(reader, "%d %d %d %d %d\n", &a, &b, &c, &d, &e)
	result := 0 
	for i:=1 ; ; i++ {
		count := 0 
		if i%a == 0 {
			count++
		}
		if i%b == 0 {
			count++
		}
		if i%c == 0 {
			count++
		}
		if i%d == 0 {
			count++
		}
		if i%e == 0 {
			count++
		}
		if count >= 3 {
			result = i
			break
		}
	}
	fmt.Fprintln(writer, result)
}