package utils

type Direction int

const (
	DirectionUp = iota
	DirectionDown
	DirectionLeft
	DirectionRight
)

type Point2D struct {
	X, Y int
}

func (p *Point2D) Move(dirs ...Direction) *Point2D {
	res := &Point2D{X: p.X, Y: p.Y}
	for _, dir := range dirs {
		switch dir {
		case DirectionUp:
			res.Y--
		case DirectionDown:
			res.Y++
		case DirectionLeft:
			res.X--
		case DirectionRight:
			res.X++
		}
	}
	return res
}
