package geonames

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdmin2Codes(t *testing.T) {
	Convey("Test Admin2Codes handler", t, func() {
		result, err := Admin2Codes()
		Convey("We should not get an error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Result should be correct", func() {
			So(result, ShouldNotBeNil)
			So(result, ShouldNotBeEmpty)
			So(result[0].Name, ShouldEqual, "Shighnan District")
		})
	})
}
