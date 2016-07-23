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
			So(result[0].Iso2Code, ShouldEqual, "AD")
		})
	})
}
