package cte

type Encoder struct {
}

func NewEncoder() *Encoder {
	this := new(Encoder)
	return this
}

func (this *Encoder) PositiveInt(value uint64) error {
	return nil
}
