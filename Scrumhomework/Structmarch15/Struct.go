package main

import "fmt"

type Students struct {
	FirstName   string
	LastName    string
	SNumber     int
	ProgramName string
	CourseList  []string
}

func main() {
	var S1 Students
	S1.FirstName = "Mark"
	S1.LastName = "cubin"
	S1.SNumber = 353727
	S1.ProgramName = "Cyber Security"
	S1.CourseList = []string{"Data Communication", "C++", "Python", "Networking"}
	fmt.Println(S1)
	fmt.Println("the student number is:", S1.SNumber)

	var S2 Students
	fmt.Print("Please provide student 2 First name:\n")
	fmt.Scanln(&S2.FirstName)
	fmt.Print("Please provide student 2 Last name:\n")
	fmt.Scanln(&S2.LastName)
	fmt.Print("Please provide student 2 SNumber:\n")
	fmt.Scanln(&S2.SNumber)
	fmt.Print("Please provide student 2 Program Name:\n")
	fmt.Scanln(&S2.ProgramName)
	fmt.Print("Please provide student 2 Couse list:\n")
	fmt.Scanln(&S2.CourseList)
	fmt.Println(S2)
}
