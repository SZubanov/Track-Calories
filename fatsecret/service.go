package fatsecret

import (
	"github.com/SZubanov/Track-Calories/helpers"
	"strconv"
)

func GetYesterdayWeight(daysWeight Month) (float64, error) {
	for _, day := range daysWeight.Day {
		dateUnix, err := strconv.Atoi(day.Date)
		if err != nil {
			return 0, err
		}

		if dateUnix == helpers.GetYesterdayUnix() {
			floatWeight, err := strconv.ParseFloat(day.Weight, 8)
			if err != nil {
				return 0, err
			}

			return floatWeight, nil
		}
	}

	return 0, nil
}
