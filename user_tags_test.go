// +build external

package geonames

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUserTags(t *testing.T) {
	Convey("Testing UserTags handler", t, func() {
		c := &Client{}
		result, err := c.UserTags()
		Convey("We should not get an error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Result should be correct", func() {
			So(result, ShouldNotBeNil)
			So(result, ShouldNotBeEmpty)
			So(result[2599253], ShouldNotBeNil)
			So(result[2599253][0], ShouldEqual, "opengeodb")
		})
	})
}
