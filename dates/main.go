package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	fmt.Println("Current date:", current)

	fmt.Println("MM-DD-YYYY :", current.Format("01-02-2006"))
	fmt.Println("YYYY-MM-DD hh:mm:ss :", current.Format("2006-01-02 15:04:05"))

	nextYear := current.AddDate(1, 0, 0)
	fmt.Println("Next year:", nextYear)

	nextMonth := current.AddDate(0, 1, 0)
	fmt.Println("Next month:", nextMonth)

	nextDay := current.AddDate(0, 0, 1)
	fmt.Println("Next day:", nextDay)

	previousYear := current.AddDate(-1, 0, 0)
	fmt.Println("Previous year:", previousYear)

	tenMoreMinutes := current.Add(10 * time.Minute)
	fmt.Println("10min into future :", tenMoreMinutes)

	tenMoreHours := current.Add(10 * time.Hour)
	fmt.Println("10h into future :", tenMoreHours)

	// difference in dates
	first := time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)
	second := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)

	difference := second.Sub(first)
	fmt.Printf("Difference in dates is %v", difference)

	// Parsing dates and times from strings
	str := "2018-08-08T11:45:26.371Z"
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Date from string:", t.String())
}
