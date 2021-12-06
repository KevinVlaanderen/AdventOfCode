import { Task } from "../types";

export const task1: Task = (data) => {
  const map = data
    .split("\n")
    .filter((line) => line !== "")
    .map((line) => {
      const data = line.match(/(\d+),(\d+) -> (\d+),(\d+)/)!;
      return {
        from: { x: parseInt(data[1], 10), y: parseInt(data[2], 10) },
        to: { x: parseInt(data[3], 10), y: parseInt(data[4], 10) },
      };
    })
    .filter((line) => line.from.x === line.to.x || line.from.y === line.to.y)
    .reduce<{ [x: number]: { [y: number]: number } }>((result, current) => {
      for (const x of range(current.from.x, current.to.x, 1)) {
        for (const y of range(current.from.y, current.to.y, 1)) {
          if (result[x] === undefined) {
            result[x] = { [y]: 1 };
          } else {
            result[x][y] = result[x][y] ? result[x][y] + 1 : 1;
          }
        }
      }
      return result;
    }, {});

  return Object.values(map)
    .flatMap((line) => Object.values(line))
    .filter((count) => count >= 2).length;
};

function range(start: number, stop: number, step: number) {
  return Array.from(
    { length: Math.abs(stop - start) / step + 1 },
    (_, i) => (start > stop ? stop : start) + i * step
  );
}
