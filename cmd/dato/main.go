package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	danish := danishDate(now)

	_, week := now.ISOWeek()
	fmt.Printf("%s, Uge %d\n", danish, week)
}

func danishDate(t time.Time) string {
	return fmt.Sprintf("%s, %02d %s '%02d",
		danishWeekday[t.Weekday()],
		t.Day(),
		danishMonth[t.Month()],
		t.Year()%100,
	)
}

var danishWeekday = map[time.Weekday]string{
	time.Monday:    "Mandag",
	time.Tuesday:   "Tirsdag",
	time.Wednesday: "Onsdag",
	time.Thursday:  "Torsdag",
	time.Friday:    "Fredag",
	time.Saturday:  "Lørdag",
	time.Sunday:    "Søndag",
}

var danishMonth = map[time.Month]string{
	time.January:   "Januar",
	time.February:  "Februar",
	time.March:     "Marts",
	time.April:     "April",
	time.May:       "Maj",
	time.June:      "Juni",
	time.July:      "Juli",
	time.August:    "August",
	time.September: "September",
	time.October:   "Oktober",
	time.November:  "November",
	time.December:  "December",
}
