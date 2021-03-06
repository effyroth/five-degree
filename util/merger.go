package util

func Inter(a *[]int, b *[]int) *[]int {
	c := []int{}

	var ahead int
	var bhead int
	for ahead != len(*a) && bhead != len(*b) {
		d := (*a)[ahead]
		e := (*b)[bhead]
		if d == e {
			c = append(c, d)
			ahead++
			bhead++

		} else if d < e {
			ahead++
		} else {
			bhead++
		}
	}
	return &c
}

func Union(a *[]int, b *[]int) *[]int {
	if len(*a) == 0 {
		return b
	}
	c := []int{}
	var ahead int
	var bhead int
	for ahead != len(*a) && bhead != len(*b) {
		d := (*a)[ahead]
		e := (*b)[bhead]
		if d == e {
			c = append(c, d)
			ahead++
			bhead++

		} else if d < e {
			c = append(c, d)
			ahead++
		} else {
			c = append(c, e)
			bhead++
		}

	}
	if ahead == len(*a) {
		for i := bhead; i < len(*b); i++ {
			c = append(c, (*b)[i])
		}
	} else if bhead == len(*b) {
		for i := ahead; i < len(*a); i++ {
			c = append(c, (*a)[i])
		}
	}
	return &c
}
