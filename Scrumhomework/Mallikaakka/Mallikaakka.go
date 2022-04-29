package main

import "fmt"

type Student struct {
	name  string
	SNo   int
	Marks int
	Grade string
}

// func Person1(name string) *Student {
// 	p := Student{name: name}
// 	p := Student{SNo: SNo}
// 	p := Student{Marks:Marks}
// 	p := Student{Grade: Grade}
// 	return &p
// }

func main() {

	//var P Student

	P1 := Student{"Manoj", 3636848, 34, "A"}

	fmt.Println(P1)

	// fmt.Println(Student{name: "Alice", SNo: 30, Marks: 45, Grade: "A"})

	// fmt.Println(Student{name: "Fred", SNo: 45, Marks: 50, Grade: "A+"})

	// fmt.Println(Person1("Jon"))

	// s := Student{name: "Sean", SNo: 50}
	// fmt.Println(s.name)

	// sp := &s
	// fmt.Println(sp.SNo)

	// sp.SNo = 51
	// fmt.Println(sp.SNo)
}
