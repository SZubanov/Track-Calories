package helpers

import "time"

func GetYesterdayUnix() int {
	startTime := time.Unix(0, 0)
	now := time.Now().AddDate(0, 0, -1)
	return int(now.Sub(startTime).Hours() / 24)
}
