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

	var L int
	var input string
	fmt.Fscan(reader, &L)
	fmt.Fscan(reader, &input)
	answer := 0
	for i, v := range input {
		num := int(v-'a')+1
		mod := 1
		for j:=0 ; j<i ; j++ {
			mod = (mod*31)%1234567891
		}
		answer = (answer+(mod*num))%1234567891
	}
	fmt.Fprintln(writer, answer)
}