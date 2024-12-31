package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	t := time.Now()
	english := t.Format("Monday, 02 January '06")

	danish := weekdays.Replace(english)
	danish = months.Replace(danish)

	_, week := t.ISOWeek()
	fmt.Println(fmt.Sprintf("%s, Uge %d", danish, week))
}

var weekdays = strings.NewReplacer(
	"Monday", "Mandag",
	"Tuesday", "Tirsdag",
	"Wednesday", "Onsdag",
	"Thursdag", "Torsdag",
	"Friday", "Fredag",
	"Saturday", "Lørdag",
	"Sunday", "Søndag",
)

var months = strings.NewReplacer(
	"January", "Januar",
	"February", "Februar",
	"March", "Marts",
	"April", "April",
	"May", "Maj",
	"June", "Juni",
	"July", "Juli",
	"August", "August",
	"September", "September",
	"October", "Oktober",
	"November", "November",
	"December", "December",
)
