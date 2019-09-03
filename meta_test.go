package aseprite

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLenAnimation(t *testing.T) {
	Convey("Given a new MetaData", t, func() {
		m := &MetaData{}
		Convey("When length of Animations is two", func() {
			m.Animations = make([]Animation, 2)
			Convey("Then result should be two", func() {
				So(m.LenAnimations(), ShouldEqual, 2)
			})
		})
		Convey("When length of Animations is zero", func() {
			m.Animations = make([]Animation, 0)
			Convey("Then result should be zero", func() {
				So(m.LenAnimations(), ShouldEqual, 0)
			})
		})
	})
}

func TestLenLayers(t *testing.T) {
	Convey("Given a new MetaData", t, func() {
		m := &MetaData{}
		Convey("When length of Layers is two", func() {
			m.Layers = make([]Layer, 2)
			Convey("Then result should be two", func() {
				So(m.LenLayers(), ShouldEqual, 2)
			})
		})
		Convey("When length of Layers is zero", func() {
			m.Layers = make([]Layer, 0)
			Convey("Then result should be zero", func() {
				So(m.LenLayers(), ShouldEqual, 0)
			})
		})
	})
}
func TestLenSlices(t *testing.T) {
	Convey("Given a new MetaData", t, func() {
		m := &MetaData{}
		Convey("When length of Slices is two", func() {
			m.Slices = make([]Slice, 2)
			Convey("Then result should be two", func() {
				So(m.LenSlices(), ShouldEqual, 2)
			})
		})
		Convey("When length of Slices is zero", func() {
			m.Slices = make([]Slice, 0)
			Convey("Then result should be zero", func() {
				So(m.LenSlices(), ShouldEqual, 0)
			})
		})
	})
}
