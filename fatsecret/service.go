package fatsecret

import (
	"github.com/SZubanov/Track-Calories/helpers"
	"strconv"
)

const WaterID = "37922"

type TotalDayEntry struct {
	Calories     float64 `json:"calories"`
	Carbohydrate float64 `json:"carbohydrate"`
	Protein      float64 `json:"protein"`
	Fat          float64 `json:"fat"`
	Fiber        float64 `json:"fiber"`
	Water        float64 `json:"water"`
}

func GetYesterdayWeight(daysWeight Month) (float64, error) {
	for _, day := range daysWeight.Day {
		dateUnix, err := strconv.Atoi(day.Date)
		if err != nil {
			return 0, err
		}
		if dateUnix == helpers.GetYesterdayUnix() {
			floatWeight := parseFloat(day.Weight)
			return floatWeight, nil
		}
	}
	return 0, nil
}

func GetTotalDayEntry(foods FoodEntries) TotalDayEntry {
	var entry TotalDayEntry
	for _, food := range foods.Foods {
		if food.ID == WaterID {
			water := parseFloat(food.Count)
			entry.Water = water
			continue
		}

		calories := parseFloat(food.Calories)
		entry.Calories += calories

		protein := parseFloat(food.Protein)
		entry.Protein += protein

		fat := parseFloat(food.Fat)
		entry.Fat += fat

		carb := parseFloat(food.Carbohydrate)
		entry.Carbohydrate += carb

		fiber := parseFloat(food.Fiber)
		entry.Fiber += fiber
	}

	return entry
}

func parseFloat(value string) float64 {
	v, err := strconv.ParseFloat(value, 8)
	if err != nil {
		v = 0
	}
	return v
}
