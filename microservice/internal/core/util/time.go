package util

import "time"

func ConvertTimeMySQLToTimeTime(layout, value string) (time.Time, error) {
	parseData, err := time.Parse(layout, value)
	if err != nil {
		return time.Time{}, err
	}
	return parseData, nil
}
