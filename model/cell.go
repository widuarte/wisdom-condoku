package model

type State string

const (
	NEWER   State = "NEWER"
	ERROR   State = "ERROR"
	DEFAULT State = "DEFAULT"
)

type Cell struct {
	State        State
	Value        string
	IsModifiable bool
}

func NewCell(value string, isModifiable bool) *Cell {
	return &Cell{
		Value:        value,
		State:        DEFAULT,
		IsModifiable: isModifiable,
	}
}

func (c *Cell) IsError() bool {
	return c.State == ERROR
}

func (c *Cell) IsNewer() bool {
	return c.State == NEWER
}

func (c *Cell) EqualsTo(cell *Cell) bool {
	return c.Value == cell.Value
}
