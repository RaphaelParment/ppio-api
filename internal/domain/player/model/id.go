package model

type Id int64

func (id Id) Int() int {
	return int(id)
}
