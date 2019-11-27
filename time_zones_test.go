// +build external

package geonames

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimeZones(t *testing.T) {
	Convey("Testing TimeZones handler", t, func() {
		c := &Client{}
		result, err := c.TimeZones()
		Convey("We should not get an error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Result should be correct", func() {
			So(result, ShouldNotBeNil)
			So(result, ShouldNotBeEmpty)
			So(result[0].CountryIso2Code, ShouldEqual, "CI")
		})
	})
}
