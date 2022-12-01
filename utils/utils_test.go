package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUtils(t *testing.T) {
	Convey("ReadFileContent", t, func(c C) {
		ec, err := ReadFileContent("../data/example.txt")
		c.So(err, ShouldBeNil)
		c.So(len(ec), ShouldEqual, 2)
	})

	Convey("ReadFileContent without filepath", t, func(c C) {
		ec, err := ReadFileContent("")
		c.So(err, ShouldNotBeNil)
		c.So(len(ec), ShouldEqual, 0)
	})

	Convey("ReadFileContent wrong filepath", t, func(c C) {
		ec, err := ReadFileContent("asdas")
		c.So(err, ShouldNotBeNil)
		c.So(len(ec), ShouldEqual, 0)
	})

	Convey("SortStringByCharacter", t, func(c C) {
		ec := SortStringByCharacter("asdas")
		c.So(ec, ShouldEqual, "aadss")
		ec = SortStringByCharacter("a")
		c.So(ec, ShouldEqual, "a")
		ec = SortStringByCharacter("")
		c.So(ec, ShouldEqual, "")
		ec = SortStringByCharacter("a132ad")
		c.So(ec, ShouldEqual, "123aad")
	})
}
