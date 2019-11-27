// +build external

package geonames

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCountryInfo(t *testing.T) {
	Convey("Testing CountryInfo handler", t, func() {
		c := &Client{}
		result, err := c.CountryInfo()
		Convey("We should not get an error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Result should be correct", func() {
			So(result, ShouldNotBeNil)
			So(result, ShouldNotBeEmpty)
			So(result[3041565].Iso2Code, ShouldEqual, "AD")
		})
	})
}
