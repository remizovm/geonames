package geonames

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestShapes(t *testing.T) {
	Convey("Testing AllShapes handler", t, func() {
		c := &Client{}
		result, err := c.Shapes()
		So(err, ShouldBeNil)
		So(result, ShouldNotBeEmpty)
	})
}
