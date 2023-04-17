package main

import (
	"DTS-Hactiv8/final-project-master/routers"
)


func main() {
	r := routers.StartApp()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
