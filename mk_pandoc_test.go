package main

import (
	"testing"
	"time"
)

func TestFillTemplate(t *testing.T) {
	template := "date={{.Tag}}~gen.~{{.Date}}~-~{{.Time}}"

	tm := time.Time{}
	meta := metadata{
		Tag:  "0.1-5-g163fcad",
		Date: formatDate(tm),
		Time: formatTime(tm),
	}

	expected := "date=0.1-5-g163fcad~gen.~01.01.1~-~00:00:00"
	actual := fillTemplate(template, meta)
	checkEquals(t, expected, actual)
}
