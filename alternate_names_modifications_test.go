package geonames

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAlternateNamesModifications(t *testing.T) {
	Convey("Test AlternateNamesModifications handler", t, func() {
		dt := time.Now()
		year := dt.Year()
		month := int(dt.Month())
		day := dt.Day()
		_, err := AlternateNamesModifications(year, month, day-1)
		Convey("Error should be nil", func() {
			So(err, ShouldBeNil)
		})
	})
}
