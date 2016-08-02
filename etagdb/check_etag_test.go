package etagdb

import (
	"testing"

	//. "github.com/smartystreets/goconvey/convey"
)

func TestCheckEtag(t *testing.T) {
	resp, err := CheckEtag("http://download.geonames.org/export/dump/readme.txt", "1f26-5386aa0b848e3")
	if err != nil {
		t.Error(err)
	}
	t.Error(resp)
}
