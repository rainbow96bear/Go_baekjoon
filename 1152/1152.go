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

	var answer int
	s,_ := reader.ReadBytes('\n')

	check := true
	for _,v := range s {
		if v != 32 && check ==true && v != 13 && v != 10 {
			check = false
			answer++
		}else if v == 32 {
			check = true
		}
	}

	
	fmt.Fprintln(writer, answer)
}