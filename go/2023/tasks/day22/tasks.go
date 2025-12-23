//go:generate go run aoc/cmd/generate-tests

package day22

import (
	"aoc/2023/tasks/day22/model"
	"aoc/framework/geometry/geo2d"
	"aoc/framework/geometry/geo2d/grid"
	"aoc/framework/geometry/geo3d"
	"aoc/framework/tasks"
	"go/types"
	"sort"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
)

// Task1 type:mock 	file:data	expected:5
// Task1 type:real 	file:day22 	expected:424
func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	bricks := parse(data)

	dropBricks(bricks)

	result.Value = lo.CountBy(bricks, func(brick *model.Brick) bool {
		if brick.Supports.Cardinality() == 0 {
			return true
		}
		for _, supported := range brick.Supports.ToSlice() {
			if supported.SupportedBy.Cardinality() == 1 {
				return false
			}
		}
		return true
	})

	return
}

// Task2 type:mock 	file:data	expected:7
// Task2 type:real 	file:day22 	expected:55483
func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	bricks := parse(data)

	dropBricks(bricks)

	result.Value = lo.SumBy(bricks, func(brick *model.Brick) int {
		fallen := mapset.NewSet[*model.Brick]()
		return countFalling(brick, bricks, &fallen)
	})

	return
}

func countFalling(brick *model.Brick, bricks []*model.Brick, fallen *mapset.Set[*model.Brick]) int {
	falling := lo.Filter(brick.Supports.ToSlice(), func(fallingBrick *model.Brick, index int) bool {
		if (*fallen).Contains(fallingBrick) {
			return false
		}
		if fallingBrick.SupportedBy.Cardinality() == 1 {
			return true
		}
		for supportingBrick := range fallingBrick.SupportedBy.Iter() {
			if supportingBrick != brick && !(*fallen).Contains(supportingBrick) {
				return false
			}
		}
		return true
	})

	result := len(falling)

	if result == 0 {
		return 0
	}

	(*fallen).Append(falling...)

	result += lo.SumBy(falling, func(fallingBrick *model.Brick) int {
		return countFalling(fallingBrick, bricks, fallen)
	})

	return result
}

func dropBricks(bricks []*model.Brick) {
	sizeX, sizeY, _ := calculateDimensions(bricks)

	heightmap := grid.NewArrayGrid[lo.Tuple2[*model.Brick, int]](sizeX, sizeY)

	sort.Sort(model.ByHeight(bricks))

	pointBuffer := make([]geo2d.Point, 0, 16)
	for _, brick := range bricks {
		pointBuffer = brick.FillXYPoints(pointBuffer)

		maxZ := 0
		for _, point := range pointBuffer {
			if highest, found := heightmap.Get(point); found && highest.B > maxZ {
				maxZ = highest.B
			}
		}
		brick.MoveToZ(maxZ + 1)
		for _, point := range pointBuffer {
			if highest, found := heightmap.Get(point); found && highest.B == maxZ && highest.A != nil {
				highest.A.Supports.Add(brick)
				brick.SupportedBy.Add(highest.A)
			}
			if ok := heightmap.Set(point, lo.Tuple2[*model.Brick, int]{A: brick, B: brick.Endpoints.B.Z}); !ok {
				panic("out of bounds")
			}
		}
	}
}

func parse(data string) []*model.Brick {
	return lo.Map(tasks.Lines(data), func(line string, index int) *model.Brick {
		parts := strings.Split(line, "~")
		return model.NewBrick(parseVoxel(parts[0]), parseVoxel(parts[1]))
	})
}

func parseVoxel(data string) geo3d.Voxel {
	parts := strings.Split(data, ",")
	var x, y, z int
	var err error
	if x, err = strconv.Atoi(parts[0]); err != nil {
		panic(err)
	}
	if y, err = strconv.Atoi(parts[1]); err != nil {
		panic(err)
	}
	if z, err = strconv.Atoi(parts[2]); err != nil {
		panic(err)
	}
	return geo3d.Voxel{X: x, Y: y, Z: z}
}

func calculateDimensions(bricks []*model.Brick) (int, int, int) {
	x := lo.MaxBy(bricks, func(a *model.Brick, b *model.Brick) bool {
		return a.Endpoints.B.X > b.Endpoints.B.X
	})
	y := lo.MaxBy(bricks, func(a *model.Brick, b *model.Brick) bool {
		return a.Endpoints.B.Y > b.Endpoints.B.Y
	})
	z := lo.MaxBy(bricks, func(a *model.Brick, b *model.Brick) bool {
		return a.Endpoints.B.Z > b.Endpoints.B.Z
	})
	return x.Endpoints.B.X + 1, y.Endpoints.B.Y + 1, z.Endpoints.B.Z + 1
}
