package geonames

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFeatureCodes(t *testing.T) {
	Convey("Testing FeatureCodes handler", t, func() {
		c := &Client{}
		result, err := c.FeatureCodes("EN")
		So(err, ShouldBeNil)
		So(result, ShouldNotBeEmpty)
	})
}
