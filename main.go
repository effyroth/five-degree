package main

import (
	"./model"
	"fmt"
)

func main() {
	fmt.Println("main start")
	p := model.NewPerson("15271919400", "effyroth")
	q := model.NewPerson("15271919400", "effyroth")
	// p.SetPerson()
	// person, _ := model.GetPerson(1)
	// fmt.Println(person)
	ids := []int{10001, 10002, 10003, 10004, 10005}
	p.SetDegreeOne(ids)
	ids = []int{10005, 10006, 10007, 10008, 10009, 10010}
	q.SetDegreeOne(ids)
	// array1 := model.GetDegreeOneUnion(p.Id, q.Id)
	// fmt.Println(array1)
	// array2 := model.GetDegreeOneInter(p.Id, q.Id)
	// fmt.Println(array2)

	// rids, _ := p.GetDegreeOne()
	// result := model.CalcDegreeTwo(&rids)
	// fmt.Println(result)
	p.SetDegreeTwo()
	degreetwo, _ := p.GetDegreeTwo()
	fmt.Println(degreetwo)
	p.IsDegreeThree(10001)

	// p.InDegreeOne(1)
	// p.InDegreeOne(6)

	// p.SetDegreeTwo(ids)
	// rids, _ = p.GetDegreeTwo()
	// fmt.Println(rids)
	// model.GetPerson(10010)
	// model.GetPerson(100000)
}
