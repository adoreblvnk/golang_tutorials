package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

// get the pointer of student so we can modify its values.
func (s *Student) setAge(age int) {
	s.age = age
}

func (s *Student) getAverageGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func (s Student) getMaxGrade() int {
	curMax := 0
	for _, v := range s.grades {
		if curMax < v {
			curMax = v
		}
	}
	return curMax
}

func main() {
	blvnk := Student{"blvnk", []int{80, 90, 50}, 17}

	// // set age
	// fmt.Println(blvnk)
	// blvnk.setAge(7)
	// fmt.Println(blvnk)

	// get average
	joe := Student{"joe", []int{50, 61}, 21}
	b_average := blvnk.getAverageGrade()
	j_average := joe.getAverageGrade()
	fmt.Println(b_average)
	fmt.Println(j_average)
	fmt.Printf("blvnk max grade: %d", blvnk.getMaxGrade())
}
