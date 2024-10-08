package main

import "fmt"

type Game struct {
	TotalPlayers int
	Name         string
	Genre        string
	Rating       float64
	Multiplayer  bool
}

type TableGame struct {
	Game
	TableNumber int
}

func main() {

	tableGame1 := TableGame{
		Game: Game{
			TotalPlayers: 5,
			Name:         "Chess",
			Genre:        "Strategy",
			Rating:       9.5,
			Multiplayer:  true,
		},
		TableNumber: 1,
	}
	tableGame2 := TableGame{
		Game: Game{
			TotalPlayers: 5,
			Name:         "NotChess",
			Genre:        "Race",
			Rating:       5,
			Multiplayer:  true,
		},
		TableNumber: 2,
	}
	tableGame3 := TableGame{
		Game: Game{
			TotalPlayers: 5,
			Name:         "DOM",
			Genre:        "Action",
			Rating:       9,
			Multiplayer:  true,
		},
		TableNumber: 3,
	}

	fmt.Printf("TableGame1: %+v\n", tableGame1)
	fmt.Printf("TableGame2: %+v\n", tableGame2)
	fmt.Printf("TableGame3: %+v\n", tableGame3)

	tableGames := make([]TableGame, 0, 4)
	tableGames = append(tableGames, tableGame1)
	fmt.Printf("TableGames: %+v\n", tableGames)
	tableGames = append(tableGames, tableGame2)
	tableGames = append(tableGames, tableGame3)
	fmt.Printf("TableGames: %+v\n", tableGames)

	//tableGameMap := map[int]TableGame{}
	//ідентичні значення
	tableGameMap := make(map[int]TableGame)
	for k, game := range tableGames {
		tableGameMap[k] = game
	}

	for key, value := range tableGameMap {
		fmt.Printf("Key: %+v, Value: %+v\n", key, value)
	}
	totalPlayers := 0
	for _, game := range tableGames {
		totalPlayers = totalPlayers + game.TotalPlayers
	}
	fmt.Println(totalPlayers)
	//_________________________________________________________________

	//// Slice
	//Slice := []int{1, 2, 3}
	//
	//Slice = append(Slice, 4, 5)
	//fmt.Println("Слайс після додавання елементів:", Slice)
	//
	//fmt.Println("Element at index 2:", Slice[2])
	//
	//newSlice := make([]int, 4)
	//fmt.Println(newSlice)
	//copy(newSlice, Slice)
	//fmt.Printf("Slice: %d\n newSlice: %d\n newSlice2: %d\n", Slice, newSlice)

	// Створюємо порожній слайс
	//slice := make([]int, 0)
	//
	//fmt.Printf("Початкова довжина: %d, місткість: %d\n", len(slice), cap(slice))
	//
	//// Додаємо елементи в слайс і слідкуємо за зміною місткості
	//for i := 1; i <= 1000; i++ {
	//	slice = append(slice, i)
	//	fmt.Printf("Довжина: %d, Місткість: %d\n", len(slice), cap(slice))
	//}

	//// Оголошуємо рядок
	//asd := "string"
	//fmt.Println(asd)
	//
	//// Перетворюємо рядок у слайс рун
	//letters := []rune(asd)
	//fmt.Printf("Runes: %c\n", letters)
	//
	//// Створюємо слайс рядків із рун
	//stringsSlice := make([]string, len(letters))
	//for i, r := range letters {
	//	stringsSlice[i] = string(r)
	//}
	//
	//// Виводимо слайс рядків
	//fmt.Printf("Strings slice: %v\n", stringsSlice)
	//
	//// Правильне виведення всього слайсу як рядка
	//fmt.Printf("Joined string: %s\n", strings.Join(stringsSlice, ""))
	//
	//// Оголошуємо новий слайс
	//dsa := make([]string, 5)
	//fmt.Println("Empty slice:", dsa)
	//
	//// Виведення типів змінних
	//fmt.Printf("Тип змінної 'dsa': %T\n", dsa)
	//fmt.Printf("Тип змінної 'letters': %T\n", letters)
	//fmt.Printf("Тип змінної 'stringsSlice': %T\n", stringsSlice)
	//
	//// Копіюємо елементи з stringsSlice в dsa
	//copy(dsa, stringsSlice)
	//fmt.Println("Copied slice:", dsa)
	//
	//// Виводимо літери з слайсу рун
	//fmt.Printf("Letters again: %c\n", letters)
	//
	//s[0] = 10
	//fmt.Println("Слайс після зміни першого елемента:", s)
	//
	//fmt.Println("Довжина слайсу:", len(s))
	//fmt.Println("Ємність слайсу:", cap(s))
	//
	//modifySlice(s)
	//fmt.Println("Слайс після зміни модифікації в функції:", s)
	//
	//s = append(s, 6, 7, 8, 9, 10)
	//fmt.Println("Слайс після розширення:", s)
	//fmt.Println("Нова ємність слайсу після розширення:", cap(s))
	//
	//// Interface
	//var i interface{} = "hello"
	//var a string
	//
	//a = i.(string)
	//fmt.Println(a)
	//
	//if a, ok := i.(string); ok {
	//	fmt.Println(a)
	//} else {
	//	fmt.Println("Value is not a string")
	//}
	//
	//// Maps
	//m := make(map[string]int)
	//
	//m["один"] = 1
	//m["два"] = 2
	//
	//fmt.Println("Значення для 'один':", m["один"])
	//
	//for key, value := range m {
	//	fmt.Printf("Ключ: %s, Значення: %d\n", key, value)
	//}
	//
	//for i := 0; i < 10; i++ {
	//	m[fmt.Sprintf("Ключ%d", i)] = i
	//}
	//fmt.Println("Кількість елементів у мапі:", len(m))
	//
	//delete(m, "два")
	//fmt.Println("Map after deletion:", m)

	//	a := 1
	//	fmt.Println(a) // 1
	//	add(a)         // 2
	//	fmt.Println(a) // 1
	//	func() {
	//		a = 3
	//		fmt.Println(a) // 3
	//	}()
	//	fmt.Println(a) // 3
	//
	//	func(a int) {
	//		a = 4
	//		fmt.Println(a) // 4
	//	}(a)
	//	fmt.Println(a) // 3
	//}
	//
	//func modifySlice(s []int) {
	//	s[2] = 100
	//}
	//
	//func add(a int) {
	//	a = 2
	//	fmt.Println(a)
}
