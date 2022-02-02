package main

import (
	"fmt"
	"github.com/SZubanov/Track-Calories/config"
	"github.com/SZubanov/Track-Calories/fatsecret"
	"github.com/SZubanov/Track-Calories/google"
	"time"
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

	monthWeight, err := fs.GetMonthWeight()
	if err != nil {
		panic(err)
	}
	weight, err := fatsecret.GetYesterdayWeight(monthWeight.Month)
	if err != nil {
		panic(err)
	}

	foodEntry, err := fs.GetFoodEntry()
	if err != nil {
		panic(err)
	}

	totalDay := fatsecret.GetTotalDayEntry(foodEntry.FoodEntries)

	yesterday := time.Now().AddDate(0, 0, -1)
	googleFormRequest := google.NewResultRequest(
		yesterday.Day(),
		int(yesterday.Month()),
		yesterday.Year(),
		weight,
		totalDay.Calories,
		totalDay.Carbohydrate,
		totalDay.Protein,
		totalDay.Fat,
		totalDay.Fiber,
		totalDay.Water,
	)

	googleFormFields := google.NewFormFields(
		conf.DayInput,
		conf.MonthInput,
		conf.YearInput,
		conf.WeightInput,
		conf.CaloriesInput,
		conf.ProteinInput,
		conf.FatInput,
		conf.CarbohydrateInput,
		conf.FiberInput,
		conf.WaterInput,
	)

	_, err = google.RequestForm(conf.FormUrl, *googleFormFields, *googleFormRequest)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done!")
}
