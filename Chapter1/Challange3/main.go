package main

import "fmt"

func main() {
	input := "selamat malam"
  
    for _, letter := range input {
        fmt.Printf("%s\n", string(letter))
    }

	data := map[string]int{
     
		" ": 1,
		"a": 4,
		"e": 1,
		"l": 2,
		"m": 3,
		"s" : 1,
		"t" : 1,
	}
	fmt.Println(data)
 
}