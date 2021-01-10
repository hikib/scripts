package main

import "testing"

func TestJournalTemplate(t *testing.T) {
	var (
		day       string = "02"
		week      string = "18"
		year      string = "1995"
		monthName string = "May"
		month     string = "05"
	)

	assertCorrectOutput := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("\ngot\n%s\n\nwant \n%s", got, want)
		}
	}

	t.Run("beautifully formatted date", func(t *testing.T) {
		got := getDate(day, monthName, year)
		want := "02.May, 1995"

		assertCorrectOutput(t, got, want)
	})

	t.Run("link to week review file", func(t *testing.T) {
		got := getWeekLink(week, year)
		want := "- [Week](../pipelines/reviews/1995-W18.md)"

		assertCorrectOutput(t, got, want)
	})

	t.Run("link to month review file", func(t *testing.T) {
		got := getMonthLink(month, year)
		want := "- [Month](../pipelines/reviews/1995-M05.md)"

		assertCorrectOutput(t, got, want)
	})

	t.Run("link to year review file", func(t *testing.T) {
		got := getYearLink(year)
		want := "- [Year](../pipelines/reviews/1995.md)"

		assertCorrectOutput(t, got, want)
	})

	t.Run("putting together the template", func(t *testing.T) {
		var (
			date      string = "02.May, 1995"
			weekLink  string = "- [Week](../pipelines/reviews/1995-W18.md)"
			monthLink string = "- [Month](../pipelines/reviews/1995-M05.md)"
			yearLink  string = "- [Year](../pipelines/reviews/1995.md)"
		)

		got := getTemplate(date, weekLink, monthLink, yearLink)
		want := `02.May, 1995

## Reviews
- [Week](../pipelines/reviews/1995-W18.md)
- [Month](../pipelines/reviews/1995-M05.md)
- [Year](../pipelines/reviews/1995.md)
`

		assertCorrectOutput(t, got, want)
	})

}
