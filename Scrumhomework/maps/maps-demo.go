package maps

import "fmt"

func MapsDemo() {
	mapofScienceMarks := make(map[string]int)
	mapofScienceMarks["Akhil"] = 100
	fmt.Println(mapofScienceMarks)
}
