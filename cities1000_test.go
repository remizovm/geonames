package geonames

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCities1000(t *testing.T) {
	Convey("Test Cities1000 handler", t, func() {
		c := &Client{}
		result, err := c.Cities1000()
		Convey("We should not get an error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Result should be correct", func() {
			So(result, ShouldNotBeNil)
			So(result, ShouldNotBeEmpty)
			So(result[3039154].ASCIIName, ShouldEqual, "El Tarter")
		})
	})
}
