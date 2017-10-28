package main

import (
	"fmt"
)

func main() {
	months := [...]string{"", "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	Q2 := months[4:7]
	fmt.Println(cap(months))
	fmt.Println(months[1:13])
	fmt.Println(Q2)
}
