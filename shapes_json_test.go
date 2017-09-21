package geonames

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestShapesJSON(t *testing.T) {
	Convey("Testing ShapesJSON handler", t, func() {
		c := &Client{}
		result, err := c.ShapesJSON()
		So(err, ShouldBeNil)
		So(result, ShouldNotBeEmpty)
	})
}
