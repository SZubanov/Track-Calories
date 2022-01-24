package main

import (
	"github.com/SZubanov/Track-Calories/fatsecret"
)

func main() {
	fs, err := fatsecret.Connect("eec54c012a134fa3a3c1cc4f2147c579", "7d6e1365dc9741fba20afec4e61dfd65")
	if err != nil {
		panic(err)
	}

	//fs.GetRequestToken()
	//fs.GetAccessToken("14273763c8a943c0b9f97a9cfe01c1f7", "88b8c303e4d049cf9d2294ebaaa9dd18", "9351096")
	//_, err = fs.RequestToken()
	//_, err = fs.GetMonthWeight()
	//_, err = fs.GetFoodEntryMonth()
	_, err = fs.GetFoodEntry()
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%+v", foods)
}

// final
//oauth_token=11faf9679618496cb5d3e05ecfd7e206&oauth_token_secret=a432f0eefc2543ca8f5877cc17e1f107
