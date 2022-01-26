package main

import (
	"fmt"
	"github.com/SZubanov/Track-Calories/config"
	"github.com/SZubanov/Track-Calories/fatsecret"
)

func main() {

	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	fs, err := fatsecret.Connect(conf.FatSecretApiKey, conf.FatSecretSecret, conf.OAuthToken, conf.OAuthTokenSecret)
	if err != nil {
		panic(err)
	}

	//monthWeight, err := fs.GetMonthWeight()
	//weight, err := fatsecret.GetYesterdayWeight(monthWeight.Month)
	//if err != nil {
	//	panic(err)
	//}

	foodEntry, err := fs.GetFoodEntry()

	fmt.Printf("%+v", foodEntry)
	//if err != nil {
	//	panic(err)
	//}

}
