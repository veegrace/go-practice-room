package calculator

//calculator interace
type calculatorApp interface {
	add() float64
	subtract() float64
	multiply() float64
	divide() float64
}

//Numbers Struct
type Numbers struct {
	X, Y float64
}

//Methods
func (n Numbers) add() float64 {
	return n.X + n.Y
}

func (n Numbers) subtract() float64 {
	return n.X - n.Y
}

func (n Numbers) multiply() float64 {
	return n.X * n.Y
}

func (n Numbers) divide() float64 {
	return n.X / n.Y
}
