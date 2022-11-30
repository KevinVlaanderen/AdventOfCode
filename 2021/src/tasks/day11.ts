import { withLines } from "../util/data";
import { range } from "../util/array";
import { applyToMatrix } from "../util/matrix";

export const task1 = withLines((data) => {
  const field = data.map((line) =>
    Array.from(line).map((item) => parseInt(item, 10))
  );

  let flashes = 0;

  for (const _ of range(1, 100)) {
    applyToMatrix(field, (x, y) => {
      flashes += tryFlash(x, y, field);
    });
    applyToMatrix(field, (x, y) => {
      if (field[y][x] > 9) {
        field[y][x] = 0;
      }
    });
  }

  return flashes;
});

export const task2 = withLines((data) => {
  const field = data.map((line) =>
    Array.from(line).map((item) => parseInt(item, 10))
  );

  let step = 1;

  while (true) {
    applyToMatrix(field, (x, y) => {
      tryFlash(x, y, field);
    });
    applyToMatrix(field, (x, y) => {
      if (field[y][x] > 9) {
        field[y][x] = 0;
      }
    });

    if (field.every((row) => row.every((col) => col === 0))) {
      return step;
    }

    step += 1;
  }
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
