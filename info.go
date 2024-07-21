package aseprite

// AnimationInfo stores essential information used
// when using the update game loop.
type AnimationInfo struct {
	CurrentAnimation *Animation

	PlaySpeed float32

	Playing           bool
	pingedOnce        bool
	animationFinished bool

	PrevFrame        int
	PrevCurrentFrame int // what?
	CurrentFrame     int
	FrameCounter     float32
}

func setupAnimationInfo() AnimationInfo {
	return AnimationInfo{
		PlaySpeed: 1.0,
	}
}

// IsPlaying returns a boolean if the provided animation is currently being
// used for the Game Loop.
func (i *AnimationInfo) IsPlaying(animation string) bool {
	return i.CurrentAnimation != nil && i.CurrentAnimation.Name == animation
}

// AnimationFinished returns a boolean if the currently used animation being
// used for the Game Loop is finished playing.
func (i *AnimationInfo) AnimationFinished() bool {
	return i.animationFinished
}

func (i *AnimationInfo) playAnimation(anim *Animation) {
	i.CurrentAnimation = anim
	i.CurrentFrame = i.CurrentAnimation.From
	i.animationFinished = false // default value
	i.pingedOnce = false        // default value
	// If the animation is reverse then start the animation from the end.
	if i.CurrentAnimation.Direction == PlayReverse {
		i.CurrentFrame = i.CurrentAnimation.To
	}
	i.FrameCounter = 0
}

func (i *AnimationInfo) advanceFrame() {
	// Reset the frame counter.
	i.FrameCounter = 0

	// Increment or decrement the current frame.
	switch i.CurrentAnimation.Direction {
	case PlayReverse:
		i.CurrentFrame--
	case PlayPingPong:
		if i.pingedOnce {
			i.CurrentFrame--
			break
		}
		fallthrough
	case PlayForward:
		fallthrough
	default:
		i.CurrentFrame++
	}
}

func (i *AnimationInfo) updateAnimationValues() {
	switch i.CurrentAnimation.Direction {
	case PlayForward:
		if i.CurrentFrame > i.CurrentAnimation.To {
			i.CurrentFrame = i.CurrentAnimation.From
			i.animationFinished = true //? return
		}
	case PlayReverse:
		if i.CurrentFrame < i.CurrentAnimation.From {
			i.CurrentFrame = i.CurrentAnimation.To
			i.animationFinished = true //? return
		}
	case PlayPingPong:
		if i.CurrentFrame > i.CurrentAnimation.To {
			i.pingedOnce = !i.pingedOnce
			i.CurrentFrame = i.CurrentAnimation.To - 1

			if i.CurrentFrame < i.CurrentAnimation.From {
				i.CurrentFrame = i.CurrentAnimation.From
				i.animationFinished = true //? return
			}
		} else if i.CurrentFrame < i.CurrentAnimation.From {
			i.pingedOnce = !i.pingedOnce
			i.CurrentFrame = i.CurrentAnimation.From + 1
			i.animationFinished = true //? return

			if i.CurrentFrame > i.CurrentAnimation.To {
				i.CurrentFrame = i.CurrentAnimation.To
			}
		}
	}
}
