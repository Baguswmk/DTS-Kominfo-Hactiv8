package main

import "fmt"

// Buatlah sebuah program go dengan :
// 1. menampilkan nilai i : 21 fmt.Printf("%T \n", i) // contoh : fmt.Printf("%v \n", i)
// 2. menampilkan tipe data dari variabel i
// 3. menampilkan tanda %
// 4. menampilkan nilai boolean j : true
// 5. menampilkan nilai boolean j : true
// 6. menampilkan unicode russia : Я (ya)
// 7. menampilkan nilai base 10 : 21
// 8. menampilkan nilai base 8 :25
// 9. menampilkan nilai base 16 : f
// 10. menampilkan nilai base 16 : F 13
// 11. menampilkan unicode karakter Я : U+042F var k float64 = 123.456;
// 12. menampilkan float : 123.456000
// 13. menampilkan float scientific : 1.234560E+02

func main() {
	var i int = 21
	var j bool = true;
	var k float64 = 123.456;

	//Nomor 1
	fmt.Printf("%v \n", i) 

	//Nomor 2
	fmt.Printf("%T \n", i)
	
	//Nomor 3
	fmt.Print("% \n")
	
	//Nomor 4
	fmt.Printf("%t \n\n", j)
	
	//Nomor 5
	fmt.Printf("%b \n", i)
	
	//Nomor 6
	fmt.Printf("%c \n", '\u042F') 

	//Nomor 7
    fmt.Printf("%d \n", i)
	
	//Nomor 8
	fmt.Printf("%o \n", i)
	
	//Nomor 9
	fmt.Printf("%x \n", 15)
	
	//Nomor 10
	fmt.Printf("%X \n", 15)

	//Nomor 11
	fmt.Printf("%U \n\n", 'Я') 

	//Nomor 12
 	fmt.Printf("%f \n", k)
	
	//Nomor 13 
 	fmt.Printf("%E \n", k) 

}