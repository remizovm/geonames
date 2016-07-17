package geonames

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLanguageCodes(t *testing.T) {
	Convey("Testing LanguageCodes handler", t, func() {
		result, err := LanguageCodes()
		Convey("We should not get an error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Result should be correct", func() {
			So(result, ShouldNotBeNil)
			So(result, ShouldNotBeEmpty)
			So(len(result), ShouldEqual, 7766)
			So(result[0].Iso3, ShouldEqual, "afa")
			So(result[len(result)-1].Iso3, ShouldEqual, "zzj")
		})
	})
}
