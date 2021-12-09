import { range, zip } from "../util";
import { withLines } from "../parsers";

type LineData = {
  from: { x: number; y: number };
  to: { x: number; y: number };
};

export const task1 = withLines((data) => {
  const lineData = extractLines(data).filter(
    (line) => line.from.x === line.to.x || line.from.y === line.to.y
  );

  return calculateResult(drawLines(lineData));
});

export const task2 = withLines((data) => {
  const lineData = extractLines(data);
  return calculateResult(drawLines(lineData));
});

function extractLines(data: string[]) {
  return data.map((line) => {
    const data = line.match(/(\d+),(\d+) -> (\d+),(\d+)/)!;
    return {
      from: { x: parseInt(data[1], 10), y: parseInt(data[2], 10) },
      to: { x: parseInt(data[3], 10), y: parseInt(data[4], 10) },
    };
  });
}

function drawLines(lineData: LineData[]) {
  return lineData.reduce<{ [x: number]: { [y: number]: number } }>(
    (result, current) => {
      const rangeX = range(current.from.x, current.to.x, 1);
      const rangeY = range(current.from.y, current.to.y, 1);

      const pairs =
        rangeX.length === rangeY.length
          ? zip(rangeX, rangeY)
          : rangeX.length > rangeY.length
          ? rangeX.map((x) => [x, rangeY[0]])
          : rangeY.map((y) => [rangeX[0], y]);

      for (const [x, y] of pairs) {
        if (result[x] === undefined) {
          result[x] = { [y]: 1 };
        } else {
          result[x][y] = result[x][y] ? result[x][y] + 1 : 1;
        }
      }
      return result;
    },
    {}
  );
}

function calculateResult(map: { [p: number]: { [p: number]: number } }) {
  return Object.values(map)
    .flatMap((line) => Object.values(line))
    .filter((count) => count >= 2).length;
}
