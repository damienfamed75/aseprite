package aseprite

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSetupAnimationInfo(t *testing.T) {
	Convey("Given no values", t, func() {
		Convey("When setupAnimationInfo is called", func() {
			a := setupAnimationInfo()
			Convey("Then PlaySpeed should equal 1.0", func() {
				So(a.PlaySpeed, ShouldEqual, 1.0)
			})
		})
	})
}

func TestIsPlaying(t *testing.T) {
	a := setupAnimationInfo()

	Convey("Given a CurrentAnimation", t, func() {
		a.CurrentAnimation = &Animation{Name: "animation"}
		Convey("When a correct animation name is provided", func() {
			res := a.IsPlaying("animation")
			Convey("Then result should be true", func() {
				So(res, ShouldBeTrue)
			})
		})
		Convey("When an incorrect animation name is provided", func() {
			res := a.IsPlaying("incorrect")
			Convey("Then result should be false", func() {
				So(res, ShouldBeFalse)
			})
		})
	})
}

func TestAnimationFinished(t *testing.T) {
	Convey("Given a new AnimationInfo", t, func() {
		a := setupAnimationInfo()
		Convey("When animationFinished is true", func() {
			a.animationFinished = true
			Convey("Then result should be true", func() {
				So(a.AnimationFinished(), ShouldBeTrue)
			})
		})
		Convey("When animationFinished is false", func() {
			a.animationFinished = false
			Convey("Then result should be false", func() {
				So(a.AnimationFinished(), ShouldBeFalse)
			})
		})
	})
}

func TestPlayAnimation(t *testing.T) {
	Convey("Given a blank AnimationInfo", t, func() {
		a := setupAnimationInfo()
		Convey("When playAnimation is provided a forward animation", func() {
			from := 1
			a.playAnimation(&Animation{
				Name: "animation",
				From: from, To: 2,
				Direction: PlayForward,
			})

			Convey("Then current frame should be equal 'from'", func() {
				So(a.CurrentFrame, ShouldEqual, from)
			})
		})
		Convey("When playAnimation is provided a reverse animation", func() {
			to := 2
			a.playAnimation(&Animation{
				Name: "animation",
				From: 1, To: to,
				Direction: PlayReverse,
			})

			Convey("Then current frame should be equal 'to'", func() {
				So(a.CurrentFrame, ShouldEqual, to)
			})
		})
		Convey("When playAnimation is provided a pingpong animation", func() {
			from := 1
			a.playAnimation(&Animation{
				Name: "animation",
				From: from, To: 2,
				Direction: PlayPingPong,
			})

			Convey("Then current frame should be equal 'from'", func() {
				So(a.CurrentFrame, ShouldEqual, from)
			})
		})
	})
}

func TestAdvanceFrame(t *testing.T) {
	a := setupAnimationInfo()
	Convey("Given a forward animation", t, func() {
		from := 1
		a.playAnimation(&Animation{
			From: from, To: 2,
			Direction: PlayForward,
		})

		So(a.CurrentFrame, ShouldEqual, from)

		Convey("When we advance frame once", func() {
			a.advanceFrame()
			Convey("Then current frame should be greater than 'from'", func() {
				So(a.CurrentFrame, ShouldBeGreaterThan, from)
			})
		})
	})
	Convey("Given a reverse animation", t, func() {
		to := 2
		a.playAnimation(&Animation{
			From: 1, To: to,
			Direction: PlayReverse,
		})

		So(a.CurrentFrame, ShouldEqual, to)

		Convey("When we advance frame once", func() {
			a.advanceFrame()
			Convey("Then current frame should be less than 'to'", func() {
				So(a.CurrentFrame, ShouldBeLessThan, to)
			})
		})
	})
	Convey("Given a pingpong animation", t, func() {
		from := 1
		to := 2
		a.playAnimation(&Animation{
			From: from, To: to,
			Direction: PlayPingPong,
		})

		So(a.CurrentFrame, ShouldEqual, from)

		Convey("When we advance frame once", func() {
			a.advanceFrame()
			Convey("Then current frame should be greater than 'from'", func() {
				So(a.CurrentFrame, ShouldBeGreaterThan, from)
			})
		})
		Convey("When animation pinged once", func() {
			a.CurrentFrame = to
			a.pingedOnce = true
			a.advanceFrame()
			Convey("Then current frame should be less than 'to'", func() {
				So(a.CurrentFrame, ShouldBeLessThan, to)
			})
		})
	})
}

func TestUpdateAnimationValues(t *testing.T) {
	Convey("Given a blank AnimationInfo", t, func() {
		a := setupAnimationInfo()
		Convey("When provided a forward animation", func() {
			from, to := 1, 2
			a.playAnimation(&Animation{
				From: from, To: to,
				Direction: PlayForward,
			})
			Convey("And the current frame is greater than 'to'", func() {
				a.CurrentFrame = to + 1
				a.updateAnimationValues()

				Convey("Then current frame should equal 'from'", func() {
					So(a.CurrentFrame, ShouldEqual, from)
				})
				Convey("And the animation should finish", func() {
					So(a.AnimationFinished(), ShouldBeTrue)
				})
			})
		})
		Convey("When provided a reverse animation", func() {
			from, to := 1, 2
			a.playAnimation(&Animation{
				From: from, To: to,
				Direction: PlayReverse,
			})
			Convey("And the current frame is less than 'from'", func() {
				a.CurrentFrame = from - 1
				a.updateAnimationValues()

				Convey("Then current frame should equal 'to'", func() {
					So(a.CurrentFrame, ShouldEqual, to)
				})
				Convey("And the animation should finish", func() {
					So(a.AnimationFinished(), ShouldBeTrue)
				})
			})
		})
		Convey("When provided a pingpong animation", func() {
			from, to := 3, 2
			a.playAnimation(&Animation{
				From: from, To: to,
				Direction: PlayPingPong,
			})
			Convey("And the current frame is less than 'from'", func() {
				a.CurrentFrame = from - 1
				a.updateAnimationValues()

				Convey("Then current frame should equal 'to'", func() {
					So(a.CurrentFrame, ShouldEqual, to)
				})
				Convey("And the animation should finish", func() {
					So(a.AnimationFinished(), ShouldBeTrue)
				})
			})
			Convey("And the current frame is greater than 'to'", func() {
				a.CurrentFrame = to + 1
				a.updateAnimationValues()

				Convey("Then current frame should equal 'to'", func() {
					So(a.CurrentFrame, ShouldEqual, to)
				})
				Convey("And the animation should finish", func() {
					So(a.AnimationFinished(), ShouldBeFalse)
				})
			})
		})
	})
}
