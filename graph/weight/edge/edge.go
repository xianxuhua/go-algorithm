package edge

type Edge struct {
	I, J int
	W    float32
}

func (e *Edge) Other(i int) int {
	if i == e.I {
		return e.J
	} else {
		return e.I
	}
}
