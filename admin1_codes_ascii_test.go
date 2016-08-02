package geonames

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdmin1CodesAscii(t *testing.T) {
	Convey("Test Admin1CodesAscii handler", t, func() {
		result, err := Admin1CodesASCII()
		Convey("We should not get an error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Result should be correct", func() {
			So(result, ShouldNotBeNil)
			So(result, ShouldNotBeEmpty)
			So(result[0].Name, ShouldEqual, "Sant Julià de Loria")
		})
	})
}
