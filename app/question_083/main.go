package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Now:", now)

	oneHourLater := now.Add(1 * time.Hour)
	fmt.Println("+1 hour:", oneHourLater)

	oneYearTwoMonthsThreeDays := now.AddDate(1, 2, 3)
	fmt.Println("+1y 2m 3d:", oneYearTwoMonthsThreeDays)

	t1 := time.Now()
	time.Sleep(100 * time.Millisecond)
	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Println("Duration (t2 - t1):", diff)
}
