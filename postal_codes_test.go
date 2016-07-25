package geonames

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPostalCodes(t *testing.T) {
	Convey("Testing PostalCodes handler", t, func() {
		result, err := PostalCodes("AD")
		Convey("We should not get an error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Result should be correct", func() {
			So(result, ShouldNotBeNil)
			So(result, ShouldNotBeEmpty)
			So(result["AD100"], ShouldNotBeNil)
			So(result["AD100"].PlaceName, ShouldEqual, "Canillo")
		})
	})
}
