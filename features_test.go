// +build external

package geonames

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFeatures(t *testing.T) {
	Convey("Test Features handler", t, func() {
		c := &Client{}
		result, err := c.Features("UA")
		Convey("We should not get an error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Result should be correct", func() {
			So(result, ShouldNotBeNil)
			So(result, ShouldNotBeEmpty)
			So(result[0].ASCIIName, ShouldEqual, "Stantsiya Krasnopartyzans'ka")
		})
	})
}
