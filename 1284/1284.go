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
		var input int
		fmt.Fscanf(reader, "%d\n", &input)

		if input == 0 {
			return
		}
		total := 1
		for input > 0 {
			num := input%10
			switch num {
				case 1 :
					total += 3
				case 0 :
					total += 5
				default :
					total += 4
			}
			input /= 10
		}
		fmt.Fprintln(writer, total)
	}
}