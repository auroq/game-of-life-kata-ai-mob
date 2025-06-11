package gameoflife

type Cell struct {
	alive bool
}

func NewCell() *Cell {
	return &Cell{alive: false}
}

func (c *Cell) IsAlive() bool {
	return c.alive
}

func (c *Cell) SetAlive(alive bool) {
	c.alive = alive
}