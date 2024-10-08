package main

import "fmt"

func main() {
	//books := map[string]int{}
	//books["book1"] = 100
	//books["book2"] = 200
	//books["book3"] = 300
	//read(books)

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	sl := []int{1, 2, 3}
	fmt.Println(sl[3])
}

//func read(books map[string]int) {
//	fmt.Println(books)
//}
