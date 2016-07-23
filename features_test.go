package geonames

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFeatures(t *testing.T) {
	Convey("Test Features handler", t, func() {
		result, err := Features("UA")
		Convey("We should not get an error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Result should be correct", func() {
			So(result, ShouldNotBeNil)
			So(result, ShouldNotBeEmpty)
			So(result[0].AsciiName, ShouldEqual, "Stantsiya Krasnopartyzans'ka")
		})
	})
}
