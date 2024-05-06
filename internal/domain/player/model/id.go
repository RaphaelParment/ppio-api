package model

type Id int64

func (id Id) AsInt() int {
	return int(id)
}

func NewUndefinedId() Id {
	return Id(0)
}
