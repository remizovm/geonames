package geonames

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAllCountries(t *testing.T) {
	Convey("Testing AllCountries handler", t, func() {
		result, err := AllCountries()
		Convey("We should not get an error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Result should be valid", func() {
			So(result, ShouldNotBeNil)
			So(result, ShouldNotBeEmpty)
		})
	})
}
