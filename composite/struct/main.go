package main

import (
	"fmt"
	"time"
)

func main() {
	type Employee struct {
		ID        int
		Name      string
		Address   string
		Dob       time.Time
		Position  string
		Salary    int
		ManagerID int
	}

	var dilbert Employee

	dilbert.Salary -= 1000
	fmt.Println(dilbert.Salary)

	position := &dilbert.Position
	*position = "Senior " + *position
	fmt.Println(*position)

	employeeOfTheMonth := &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	fmt.Println(employeeOfTheMonth.Position)
}
