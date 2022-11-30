import { withBlocks } from "../util/data";
import { range } from "../util/array";

enum Direction {
  UP,
  LEFT,
}

type Grid = Record<number, Record<number, boolean>>;

export const task1 = withBlocks((data) => {
  const coordinates = parseCoordinates(data[0]);
  const instructions = parseInstructions(data[1]);

  const grid = initializeGrid(coordinates);

  applyFoldInstruction(grid, ...instructions[0]);

  return countDots(grid);
});

export const task2 = withBlocks((data) => {
  const coordinates = parseCoordinates(data[0]);
  const instructions = parseInstructions(data[1]);

  const grid = initializeGrid(coordinates);

  instructions.forEach((instruction) =>
    applyFoldInstruction(grid, ...instruction)
  );

  const maxX = Math.max(
    ...Object.values(grid).flatMap((column) =>
      Object.keys(column).map((x) => parseInt(x, 10))
    )
  );
  const maxY = Math.max(...Object.keys(grid).map((y) => parseInt(y, 10)));

  const printGrid = range(0, maxY).map((y) =>
    range(0, maxX).map((x) => (grid[y] && grid[y][x] ? "#" : "."))
  );

  console.log(printGrid.map((line) => line.join("")));

  return "FGKCKBZG";
});

function parseCoordinates(data: string): Array<[number, number]> {
  return data.split("\n").map((line) => {
    const parts = line.split(",");
    return [parseInt(parts[0], 10), parseInt(parts[1], 10)];
  });
}

function parseInstructions(data: string): Array<[Direction, number]> {
  return data.split("\n").map((line) => {
    const parts = line.match(/fold along (.+)=(\d+)/)!;
    return [
      parts[1] === "x" ? Direction.LEFT : Direction.UP,
      parseInt(parts[2], 10),
    ];
  });
}

function initializeGrid(coordinates: Array<[number, number]>) {
  return coordinates.reduce<Grid>(
    (result, [x, y]) => ({
      ...result,
      [y]: {
        ...result[y],
        [x]: true,
      },
    }),
    {}
  );
}

function applyFoldInstruction(
  grid: Grid,
  direction: Direction,
  foldLine: number
) {
  if (direction === Direction.UP) {
    foldUp(grid, foldLine);
  } else {
    foldLeft(grid, foldLine);
  }
}

function foldUp(grid: Grid, foldLine: number) {
  for (const [indexY, row] of Object.entries(grid)) {
    const y = parseInt(indexY, 10);

    if (y > foldLine) {
      const newY = foldLine - (y - foldLine);

      for (const indexX of Object.keys(row)) {
        const x = parseInt(indexX, 10);

        grid[newY] = { ...grid[newY], [x]: true };
        delete grid[y][x];
      }
    }
  }
}

function foldLeft(grid: Grid, foldLine: number) {
  for (const [indexY, row] of Object.entries(grid)) {
    const y = parseInt(indexY, 10);

    for (const indexX of Object.keys(row)) {
      const x = parseInt(indexX, 10);

      if (x > foldLine) {
        const newX = foldLine - (x - foldLine);

        grid[y] = { ...grid[y], [newX]: true };
        delete grid[y][x];
      }
    }
  }
}

function countDots(grid: Grid) {
  return Object.values(grid).reduce(
    (result, column) =>
      result + Object.values(column).filter((value) => value === true).length,
    0
  );
}
