package main

import "fmt"

func main() {
	// creating a map
	// var mp map[string]int = map[string]int{
	// 	"apple": 5,
	// 	"pear":  1,
	// 	"peach": 4,
	// }

	// creating a map
	mp := make(map[string]int)
	mp = map[string]int{
		"apple": 5,
		"pear":  1,
		"peach": 4,
	}
	mp["apple"] = 900
	mp["blvnk"] = 314  // creating a key value
	delete(mp, "pear") // delete a key

	// check key existence
	// val, ok := mp["apple"] // 900, true
	val, ok := mp["adore"] // 0, false
	fmt.Println(val, ok)

	fmt.Println(mp)

}
