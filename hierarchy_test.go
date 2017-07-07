package geonames

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHierarchy(t *testing.T) {
	Convey("Testing Hierarchy handler", t, func() {
		c := &Client{}
		result, err := c.Hierarchy()
		Convey("We should not get an error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Result should be correct", func() {
			So(result, ShouldNotBeNil)
			So(result, ShouldNotBeEmpty)
			So(result[6295630], ShouldNotBeNil)
			So(result[6295630][0].ChildID, ShouldEqual, 6255146)
		})
	})
}
