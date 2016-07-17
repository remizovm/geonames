package geonames

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimeZones(t *testing.T) {
	Convey("Testing TimeZones handler", t, func() {
		result, err := TimeZones()
		Convey("We should not get an error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Result should be correct", func() {
			So(result, ShouldNotBeNil)
			So(result, ShouldNotBeEmpty)
			So(len(result), ShouldEqual, 421)
			So(result[0].CountryIso2Code, ShouldEqual, "CI")
			So(result[len(result)-1].CountryIso2Code, ShouldEqual, "WF")
		})
	})
}
