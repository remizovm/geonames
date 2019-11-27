// +build external

package geonames

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLanguageCodes(t *testing.T) {
	Convey("Testing LanguageCodes handler", t, func() {
		c := &Client{}
		result, err := c.LanguageCodes()
		Convey("We should not get an error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Result should be correct", func() {
			So(result, ShouldNotBeNil)
			So(result, ShouldNotBeEmpty)
			So(result[0].Iso3, ShouldEqual, "afa")
		})
	})
}
