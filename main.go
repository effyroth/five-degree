package main

import (
	"./model"
	// "fmt"
)

func main() {
	// fmt.Println("main start")
	// p := model.NewPerson("133", "effyroth")
	// p.SetPerson()
	// person, _ := model.GetPerson(1)
	// fmt.Println(person)
	// ids := []int{1, 2, 3, 4, 5}
	// p.SetDegreeOne(ids)
	// rids, _ := p.GetDegreeOne()
	// fmt.Println(rids)
	model.GetPerson(10010)
	model.GetPerson(100000)
}
