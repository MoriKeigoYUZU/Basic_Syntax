// Code generated by "stringer -type DayOfWeek main.go"; DO NOT EDIT

package main

import "fmt"

const _DayOfWeek_name = "SundayMondayTuesdayWednesdayThursdayFridaySaturday"

var _DayOfWeek_index = [...]uint8{0, 6, 12, 19, 28, 36, 42, 50}

func (i DayOfWeek) String() string {
	if i < 0 || i >= DayOfWeek(len(_DayOfWeek_index)-1) {
		return fmt.Sprintf("DayOfWeek(%d)", i)
	}
	return _DayOfWeek_name[_DayOfWeek_index[i]:_DayOfWeek_index[i+1]]
}