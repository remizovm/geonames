// +build external

package geonames

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAlternateNamesDeletes(t *testing.T) {
	Convey("Test AlternateNamesDeletes handler", t, func() {
		c := &Client{}
		dt := time.Now()
		year := dt.Year()
		month := int(dt.Month())
		day := dt.Day()
		_, err := c.AlternateNamesDeletes(year, month, day-1)
		Convey("Error should be nil", func() {
			So(err, ShouldBeNil)
		})
	})
}
