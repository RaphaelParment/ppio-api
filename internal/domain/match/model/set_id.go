package model

type SetId int64

func NewUndefinedSetId() SetId {
	return SetId(0)
}

func (id SetId) AsInt() int {
	return int(id)
}
