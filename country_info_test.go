package geonames

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCountryInfo(t *testing.T) {
	Convey("Testing CountryInfo handler", t, func() {
		result, err := CountryInfo()
		Convey("We should not get an error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Result should be correct", func() {
			So(result, ShouldNotBeNil)
			So(result, ShouldNotBeEmpty)
			So(len(result), ShouldEqual, 252)
			So(result[0].Iso2Code, ShouldEqual, "AD")
			So(result[len(result)-1].Iso2Code, ShouldEqual, "AN")
		})
	})
}
