package mover

import "2022/src/tasks/day5/model"

type CrateMover9000 struct {
	Storage *model.Storage
}

func (c *CrateMover9000) Move(n int, from int, to int) {
	for i := 0; i < n; i++ {
		crate := c.Storage.Stacks[from].Remove()
		c.Storage.Stacks[to].Add(crate)
	}
}
