package main

import (
	"testing"
)

func TestLink(t *testing.T) {
	tests := []struct {
		name    string
		link    Link
		hasLink string
	}{
		{
			name:    "Week",
			link:    Week{year: Year{2043}, weekNum: 5},
			hasLink: "- [Week](../pipelines/reviews/2043-W05.md)",
		},
		{
			name:    "Month",
			link:    Month{year: Year{2043}, month: 3},
			hasLink: "- [Month](../pipelines/reviews/2043-M03.md)",
		},
		{
			name:    "Year",
			link:    Year{yearNum: 2020},
			hasLink: "- [Year](../pipelines/reviews/2020.md)",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.link.MarkdownLink()
			if got != test.hasLink {
				t.Errorf("%#v\ngot\n%s\n\nwant \n%s", test.link, got, test.hasLink)
			}

		})
	}
}
