package model

import (
	"../util"
	// "fmt"
)

func GetDegreeOneUnion(id1 int, id2 int) *[]int {
	person1, _ := GetPerson(id1)
	person2, _ := GetPerson(id2)
	array1, _ := person1.GetDegreeOne()
	array2, _ := person2.GetDegreeOne()
	return util.Union(&array1, &array2)

}

func GetDegreeOneInter(id1 int, id2 int) *[]int {
	person1, _ := GetPerson(id1)
	person2, _ := GetPerson(id2)
	array1, _ := person1.GetDegreeOne()
	array2, _ := person2.GetDegreeOne()
	return util.Inter(&array1, &array2)

}

func CalcDegreeTwo(ids *[]int) *[]int {
	result := []int{}
	for _, id := range *ids {
		person, _ := GetPerson(id)
		array, _ := person.GetDegreeOne()
		result = *util.Union(&array, &result)
		// fmt.Println(result, array)
	}
	return &result
}
