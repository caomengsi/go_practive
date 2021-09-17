package main

import "fmt"

func GetName()(firstName, middleName, lastName string)  {
	return "may", "and", "last"
}
type Integer int

// 传入指针
// func (a *Integer) Add(b Integer) {
//   *a += b 
// }

// 传入值
func (a Integer) Add(b Integer) {
	a += b 
}


func main()  {
	var a Integer = 1
	a.Add(2)
	fmt.Println("hello world!", a)
	// _, _, a := GetName()
	// fmt.Println(a)
	// const (
	// 	b = 1 << iota
	// 	c
	// 	d
	// )
	// fmt.Println(b, c, d)
	// v1 := 1 == 2
	// fmt.Println(v1)

	// var myArray [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}
	// var sliceArray []int = myArray[:5]
	// for _, v := range myArray {
	// 	fmt.Println(v, "  ")
	// }
	// for _, v := range sliceArray {
	// 	fmt.Println(v, "  ")
	// }

	// var person map[string]int
	// person = make(map[string]int)
	// person["123"]= 345
	// fmt.Println(person)
	// val, ok := person["123"]
	// if ok {
	// 	fmt.Println("ok", val)
	// }

	// 值传递

}