package mover

import "2022/src/tasks/day5/model"

type CrateMover9001 struct {
	Storage *model.Storage
}

func (c *CrateMover9001) Move(n int, from int, to int) {
	crates := c.Storage.Stacks[from].Crates[len(c.Storage.Stacks[from].Crates)-n:]
	for i := 0; i < n; i++ {
		c.Storage.Stacks[from].Remove()
		c.Storage.Stacks[to].Add(crates[i])
	}
}
