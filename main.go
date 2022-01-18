package main

import (
	"fmt"
	"github.com/SZubanov/Track-Calories/fatsecret"
)

func main() {
	fs, err := fatsecret.Connect("eec54c012a134fa3a3c1cc4f2147c579", "7d6e1365dc9741fba20afec4e61dfd65")
	if err != nil {
		panic(err)
	}

	foods, err := fs.GetRequestToken()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", foods)
}
