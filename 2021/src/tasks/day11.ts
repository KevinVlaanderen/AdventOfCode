import { withLines } from "../parsers";
import { range } from "../util";

export const task1 = withLines((data) => {
  const field = data.map((line) =>
    Array.from(line).map((item) => parseInt(item, 10))
  );

  let flashes = 0;

  for (const _ in range(1, 100)) {
    for (const [y, row] of field.entries()) {
      for (const [x] of row.entries()) {
        flashes += tryFlash(x, y, field);
      }
    }

    for (const [y, row] of field.entries()) {
      for (const [x] of row.entries()) {
        if (field[y][x] > 9) {
          field[y][x] = 0;
        }
      }
    }
  }

  return flashes;
});

function tryFlash(x: number, y: number, field: number[][]): number {
  if (field[y][x] === 10) return 0;

  field[y][x] += 1;

  if (field[y][x] <= 9) return 0;

  const neighbours: Array<[number, number]> = [
    [x - 1, y - 1],
    [x, y - 1],
    [x + 1, y - 1],
    [x - 1, y],
    [x + 1, y],
    [x - 1, y + 1],
    [x, y + 1],
    [x + 1, y + 1],
  ];

  return neighbours.reduce(
    (result, neighbour) =>
      validCoordinate(neighbour[0], neighbour[1], field)
        ? result + tryFlash(...neighbour, field)
        : result,
    1
  );
}

function validCoordinate(x: number, y: number, field: number[][]) {
  return y in range(0, field.length - 1) && x in range(0, field[y].length - 1);
}
