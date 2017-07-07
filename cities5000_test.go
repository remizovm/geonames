package geonames

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCities5000(t *testing.T) {
	Convey("Test Cities5000 handler", t, func() {
		c := &Client{}
		result, err := c.Cities5000()
		Convey("We should not get an error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Result should be correct", func() {
			So(result, ShouldNotBeNil)
			So(result, ShouldNotBeEmpty)
			So(result[3039163].ASCIIName, ShouldEqual, "Sant Julia de Loria")
		})
	})
}
