package processor

import (
	"anagram/entity"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestProcess(t *testing.T) {
	Convey("Process - happy path", t, func(c C) {
		contentList := []entity.Content{
			{Words: []string{"a", "b"}},
			{Words: []string{"asd", "dsa"}},
			{Words: []string{"eqwe", "qwee", "eeqw"}},
			{Words: []string{"zzzzz"}},
		}
		res := Processor(contentList)
		c.So(len(res), ShouldEqual, 2)
		c.So(res, ShouldContain, "asd,dsa")
		c.So(res, ShouldContain, "eqwe,qwee,eeqw")
	})

	Convey("Process - min lenght", t, func(c C) {
		contentList := []entity.Content{
			{Words: []string{"a", "b"}},
		}
		res := Processor(contentList)
		c.So(len(res), ShouldEqual, 0)
	})

	Convey("Process - no anagram", t, func(c C) {
		contentList := []entity.Content{
			{Words: []string{"zzzzz"}},
		}
		res := Processor(contentList)
		c.So(len(res), ShouldEqual, 0)
	})

	Convey("Process - 1 anagram", t, func(c C) {
		contentList := []entity.Content{
			{Words: []string{"zazazaza", "azazazaz"}},
		}
		res := Processor(contentList)
		c.So(len(res), ShouldEqual, 1)
	})
}
