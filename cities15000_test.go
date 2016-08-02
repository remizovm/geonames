package geonames

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCities15000(t *testing.T) {
	Convey("Test Cities15000 handler", t, func() {
		result, err := Cities15000()
		Convey("We should not get an error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Result should be correct", func() {
			So(result, ShouldNotBeNil)
			So(result, ShouldNotBeEmpty)
			So(result[3040051].ASCIIName, ShouldEqual, "les Escaldes")
		})
	})
}
