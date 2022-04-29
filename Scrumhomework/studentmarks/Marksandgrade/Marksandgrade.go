package Marksandgrade

import "fmt"

type TotalMarks struct {
	Math     int
	Computer int
	Heath    int
	//Total int
	Name string
	SNo  int
}

func Details() {
	a := TotalMarks{Math: 40, Computer: 50, Heath: 30, Name: "mam", SNo: 343}
	fmt.Println(a)
	Total := a.Computer + a.Math + a.Heath
	fmt.Println("Total Marks:", Total)

	avg := Total / 3
	fmt.Println("Average of student:", avg)
	if avg >= 40 {
		fmt.Println("Grade: A")
	} else if avg >= 35 && avg < 40 {
		fmt.Println("Grade: B")
	} else if avg >= 28 && avg < 35 {
		fmt.Println("Grade: C")
	} else if avg >= 25 && avg < 28 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}

}

// func Grade() {
// 	Details()
// 	var t TotalMarks
// 	var Total int = t.Math + t.Computer + t.Heath
// 	fmt.Println(Total)
// 	// Total = Math + Computer + Health
// 	avg := Total / 3
// 	if avg >= 90 {
// 		fmt.Println("Grade: A")
// 	} else if avg >= 80 && avg < 90 {
// 		fmt.Println("Grade: B")
// 	} else if avg >= 70 && avg < 80 {
// 		fmt.Println("Grade: C")
// 	} else if avg >= 60 && avg < 70 {
// 		fmt.Println("Grade: D")
// 	} else {
// 		fmt.Println("Grade: F")
// 	}
// }
