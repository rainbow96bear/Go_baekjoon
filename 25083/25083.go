package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	writer := bufio.NewWriter(os.Stdout)
	
	defer writer.Flush()

	fmt.Fprintln(writer,`         ,r'"7`)
	fmt.Fprintln(writer,"r`-_   ,'  ,/")
	fmt.Fprintln(writer,` \. ". L_r'`)
	fmt.Fprintln(writer,"   `~\\/")
	fmt.Fprintln(writer,"      |")
	fmt.Fprintln(writer,"      |")

}