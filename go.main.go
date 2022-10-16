package main

import "fmt"

func main() {

	var f = [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	} // true

	repeated := false

	for x, i := range f {
		//fmt.Println(x, i)
		for y := 0; y < x; y++ {
			var c9 []int
			for _, val := range i {
				for _, items := range c9 {
					if items == val {
						//fmt.Println("we are adding a repeated character to the temp slice")
						repeated = true
					}
				}
				c9 = append(c9, val)
			}
		}
	}

	var c8 []int

	var c7 []int

	for g := 0; g < 9; g++ {
		for _, x := range f {
			c8 = append(c8, x[g])
		}
	}

	index := 0
	for i, x := range c8 {
		index = i + 1
		c7 = append(c7, x)
		if index%9 == 0 {
			//fmt.Println(c7)
			var c6 []int
			for _, k := range c7 {
				for _, h := range c6 {
					if h == k {
						repeated = true
					}
				}
				c6 = append(c6, k)

			}
			c7 = nil
			//fmt.Println("------------")

		}

	}

	truth := test(f)

	if truth == true {
		repeated = true
	} else if truth == false {
		repeated = false
	}

	//{5, 3, 4, 	6, 7, 8, 	9, 1, 2},
	//{6, 7, 2, 	1, 9, 0, 	3, 4, 8},
	//{1, 0, 0, 	3, 4, 2, 	5, 6, 0},
	//{8, 5, 9, 	7, 6, 1, 	0, 2, 0},
	//{4, 2, 6, 	8, 5, 3, 	7, 9, 1},
	//{7, 1, 3, 	9, 2, 4, 	8, 5, 6},
	//{9, 0, 1, 	5, 3, 7, 	2, 1, 4},
	//{2, 8, 7, 	4, 1, 9, 	6, 3, 5},
	//{3, 0, 0, 	4, 8, 1, 	1, 7, 9},

	if repeated == true {
		repeated = false
	} else if repeated == false {
		repeated = true
	}

	fmt.Println(repeated)

}

func test(m [][]int) bool {

	board := m
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}
