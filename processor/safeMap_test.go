package processor

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSafeMap(t *testing.T) {
	Convey("SafeMap - init", t, func(c C) {
		safeMap := New()
		c.So(len(safeMap.safeMap), ShouldEqual, 0)
		safeMap.add("k", "v")
		safeMap = New()
		c.So(len(safeMap.safeMap), ShouldEqual, 1)
	})

	Convey("SafeMap - add", t, func(c C) {
		safeMap := New()
		safeMap.add("key", "value")
		c.So(len(safeMap.safeMap), ShouldEqual, 2)
	})

	Convey("SafeMap - get", t, func(c C) {
		safeMap := New()
		r := safeMap.get("key")
		c.So(r, ShouldEqual, "")
		safeMap.add("key", "value2")
		r = safeMap.get("key")
		c.So(r, ShouldEqual, "value,value2")
		r = safeMap.get("key-no-exist")
		c.So(r, ShouldEqual, "")
	})

	Convey("SafeMap - flush", t, func(c C) {
		safeMap := New()
		safeMap.flush()
		c.So(len(safeMap.safeMap), ShouldEqual, 0)
	})
}
