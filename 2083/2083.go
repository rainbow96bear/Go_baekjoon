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
		var name string
		var age, weight int

		fmt.Fscanf(reader, "%s %d %d\n", &name, &age, &weight)
		if name == "#" && age == 0 && weight == 0 {
			return
		}
		if age > 17 || weight >= 80 {
			fmt.Fprintf(writer, "%s %s\n", name, "Senior")
		}else {
			fmt.Fprintf(writer, "%s %s\n", name, "Junior")
		}
	}
}