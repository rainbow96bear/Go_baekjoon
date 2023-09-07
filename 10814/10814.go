package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)
type User struct{
	age int
	name string
	num int
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)
	UserList := make([]User, N)
	for i:=0 ; i<N ; i++{
		fmt.Fscan(reader, &UserList[i].age, &UserList[i].name)
		UserList[i].num=i
	}
	sort.Slice(UserList, func(i, j int)bool{
		if UserList[i].age == UserList[j].age{
			return UserList[i].num < UserList[j].num
		}
		return UserList[i].age < UserList[j].age
	})
	for i, _ := range UserList{
		fmt.Fprintln(writer, UserList[i].age, UserList[i].name)
	} 
}