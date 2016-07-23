package geonames

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDeletes(t *testing.T) {
	Convey("Test Deletes handler", t, func() {
		dt := time.Now()
		year := dt.Year()
		month := int(dt.Month())
		day := dt.Day()
		_, err := Deletes(year, month, day-1)
		Convey("Error should be nil", func() {
			So(err, ShouldBeNil)
		})
	})
}
