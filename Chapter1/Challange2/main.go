package main

import "fmt"

func main() {
	const char = "САШАРВО"

	for i := 0; i < 5; i++{
		fmt.Println("Nilai i : ", i)
	}
	for j := 0; j <= 10; j++ {
		if j == 5 {
			continue
		}
		fmt.Println("Nilai j : ", j) 
		
		if j == 4 {
			for index, runeValue := range char {
				index = index * 2
				fmt.Printf("character %#U starts at byte position %d\n", runeValue, index)
			}	
		}
	}
}

