package lazy

type Duplex[T any, U any] struct {
	versionLeft  uint64
	versionRight uint64
	left2right   func(T) U
	right2left   func(U) T
	left         T
	right        U
}

func NewDuplex[T any, U any](
	left T, right U,
	versionLeft uint64, versionRight uint64,
	left2right func(T) U, right2left func(U) T,
) Duplex[T, U] {
	return Duplex[T, U]{
		versionLeft:  versionLeft,
		versionRight: versionRight,
		left2right:   left2right,
		right2left:   right2left,
		left:         left,
		right:        right,
	}
}

func NewDuplexFromLeft[T any, U any](left T, left2right func(T) U, right2left func(U) T) Duplex[T, U] {
	return Duplex[T, U]{
		versionLeft:  1,
		versionRight: 0,
		left2right:   left2right,
		right2left:   right2left,
		left:         left,
	}
}

func NewDuplexFromRight[T any, U any](right U, left2right func(T) U, right2left func(U) T) Duplex[T, U] {
	return Duplex[T, U]{
		versionLeft:  0,
		versionRight: 1,
		left2right:   left2right,
		right2left:   right2left,
		right:        right,
	}
}

func (d *Duplex[T, U]) Left() T {
	if d.versionLeft >= d.versionRight {
		return d.left
	}

	d.left = d.right2left(d.right)
	d.versionLeft = d.versionRight
	return d.left
}

func (d *Duplex[T, U]) Right() U {
	if d.versionRight >= d.versionLeft {
		return d.right
	}

	d.right = d.left2right(d.left)
	d.versionRight = d.versionLeft
	return d.right
}

func (d *Duplex[T, U]) SetLeft(left T) {
	d.left = left
	d.versionLeft++
}

func (d *Duplex[T, U]) SetRight(right U) {
	d.right = right
	d.versionRight++
}
