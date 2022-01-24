package main

import (
	"github.com/SZubanov/Track-Calories/fatsecret"
)

func main() {
	fs, err := fatsecret.Connect("eec54c012a134fa3a3c1cc4f2147c579", "7d6e1365dc9741fba20afec4e61dfd65")
	if err != nil {
		panic(err)
	}

	//_, err = fs.RequestToken()
	_, err = fs.AuthToken()
	//_, err = fs.GetAuth("73909453")
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%+v", foods)
}

// final
//oauth_token=11faf9679618496cb5d3e05ecfd7e206&oauth_token_secret=a432f0eefc2543ca
//8f5877cc17e1f107
