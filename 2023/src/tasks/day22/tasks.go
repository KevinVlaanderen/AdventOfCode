package day22

import (
	"2023/src/framework"
	"2023/src/framework/geometry"
	"2023/src/framework/geometry/grid"
	"2023/src/tasks/day22/model"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
	"github.com/tidwall/pinhole"
	"go/types"
	"golang.org/x/exp/rand"
	"golang.org/x/image/colornames"
	"sort"
	"strconv"
	"strings"
)

func Task1(data string, _ types.Nil) (result framework.Result[int]) {
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

func Task2(data string, _ types.Nil) (result framework.Result[int]) {
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

	//seed := uint64(time.Now().Second())
	//rand.Seed(seed)
	//drawBricks(bricks, sizeX, sizeY, sizeZ, "before")

	heightmap := grid.NewGrid[lo.Tuple2[*model.Brick, int]](sizeX, sizeY)

	sort.Sort(model.ByHeight(bricks))

	pointBuffer := make([]geometry.Point, 0, 16)
	for _, brick := range bricks {
		pointBuffer = brick.FillXYPoints(pointBuffer)

		maxZ := 0
		for _, point := range pointBuffer {
			if highest, found := heightmap.Get(&point); found && highest.B > maxZ {
				maxZ = highest.B
			}
		}
		brick.MoveToZ(maxZ + 1)
		for _, point := range pointBuffer {
			if highest, found := heightmap.Get(&point); found && highest.B == maxZ && highest.A != nil {
				highest.A.Supports.Add(brick)
				brick.SupportedBy.Add(highest.A)
			}
			if err := heightmap.Set(&point, lo.Tuple2[*model.Brick, int]{A: brick, B: brick.Endpoints.B.Z}); err != nil {
				panic(err)
			}
		}
	}

	//rand.Seed(seed)
	//sizeX, sizeY, sizeZ = calculateDimensions(bricks)
	//drawBricks(bricks, sizeX, sizeY, sizeZ, "after")
}

func parse(data string) []*model.Brick {
	return lo.Map(framework.Lines(data), func(line string, index int) *model.Brick {
		parts := strings.Split(line, "~")
		return model.NewBrick(parseVoxel(parts[0]), parseVoxel(parts[1]))
	})
}

func parseVoxel(data string) geometry.Voxel {
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
	return geometry.Voxel{X: x, Y: y, Z: z}
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

func drawBricks(bricks []*model.Brick, x, y, z int, filename string) {
	sizeX, sizeY, sizeZ := 1.0/float64(x), 1.0/float64(y), 1.0/float64(z)
	size := max(sizeX, sizeY, sizeZ)

	p := pinhole.New()
	for _, brick := range bricks {
		p.Begin()
		p.DrawCube(
			float64(brick.Endpoints.A.X)*size,
			float64(brick.Endpoints.A.Z)*size,
			float64(brick.Endpoints.A.Y)*size,
			float64(brick.Endpoints.B.X+1)*size,
			float64(brick.Endpoints.B.Z+1)*size,
			float64(brick.Endpoints.B.Y+1)*size)
		p.Colorize(colornames.Map[colornames.Names[rand.Int()%len(colornames.Names)]])
		p.End()
	}
	p.Center()
	p.Rotate(-0.1, 0.5, 0)
	p.Translate(0, 0, 1)
	_ = p.SavePNG(filename+".png", 2000, 2000, nil)
}
