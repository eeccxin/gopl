package main

import (
	"os"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func updatePosition(emp *Employee, newPosition string) {
	emp.Position = newPosition
}

func add1(r rune) rune { return r + 1 }

func tempDirs() []string {
	var ints = []string{"1", "2", "3"}
	return ints
}

func main() {

	//var dilbert Employee
	//position := &dilbert.Position
	//*position = "Senior " + *position // promoted, for outsourcing to Elbonia
	//fmt.Printf("%+v", dilbert)
	//print(dilbert.Position, (&dilbert).Position)

	//dilbert := &Employee{
	//	0,
	//	"Dilbert",
	//	"",
	//	time.Time{},
	//	"Engineer",
	//	0,
	//	0,
	//}
	//updatePosition(dilbert, "Senior Engineer")
	//fmt.Println(dilbert.Position) // 输出 "Senior Engineer"
	//fmt.Printf("%+v", dilbert)    // "false"
	//
	//fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	//fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	//fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"
	//fmt.Println(strings.Map(add1, "0"))        // "Benjy"
	////r := rune("HAL-9000")
	////println(r)
	////str := "Hello, 世界"
	////for _, r := range str {
	////	fmt.Printf("%c %U\n", r, r)
	////}
	//r := "世"
	//str := string(r)
	//fmt.Printf("%T", r)
	//fmt.Println(str) // 输出 "世"

	var rmdirs []func()
	for _, d := range tempDirs() {
		dir := d // NOTE: necessary!
		//os.MkdirAll(dir, 0755) // creates parent directories too
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}
	// ...do some work…
	for _, rmdir := range rmdirs {
		rmdir() // clean up
	}
}
