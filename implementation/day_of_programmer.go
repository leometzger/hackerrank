package main

import "time"

var resultFormat string = "02.01.2006"
var programmerDayOfYear int = 256

// https://www.hackerrank.com/challenges/day-of-the-programmer/problem?isFullScreen=true
func DayOfProgrammer(year int32) string {
	isGregorianCalendar := year >= 1919
	isLeapYear := year%4 == 0
	hasSpecialCondition := isLeapYear && year%100 == 0

	isTransitionYear := year == 1918

	if isTransitionYear {
		return calculateDate(year, programmerDayOfYear+13)
	}

	if hasSpecialCondition && !isGregorianCalendar {
		return calculateDate(year, programmerDayOfYear-1)
	}

	return calculateDate(year, programmerDayOfYear)
}

func calculateDate(year int32, daysToAdd int) string {
	t := time.Date(int(year), time.January, 0, 12, 0, 0, 0, time.Local)

	return t.Add(time.Duration(daysToAdd) * time.Hour * 24).Format(resultFormat)
}
