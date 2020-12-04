package main

import "fmt"

type contact struct {
	Conuntry, Name, Ability, Salary string
	Age                             int
	Task                            task
}

type task struct {
	Taskname, Deadline string
}

func main() {

	Alldata := contact{Conuntry: "Uzbekistan", Name: "MarCuz", Ability: "Golang", Age: 25, Salary: "$1000", Task: task{Taskname: "Micro Service", Deadline: "12.12.2020"}}
	fmt.Println(Alldata)
}
