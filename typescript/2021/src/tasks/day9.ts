import { withLines } from "../parsers";
import { isDefined } from "../util";

export const task1 = withLines((data) => {
  const input = data.map((line) =>
    Array.from(line).map((char) => parseInt(char, 10))
  );

  return input
    .flatMap((line, y) =>
      line.flatMap((number, x) =>
        [
          [x - 1, y],
          [x, y - 1],
          [x + 1, y],
          [x, y + 1],
        ]
          .filter(
            (check) =>
              check[0] >= 0 &&
              check[1] >= 0 &&
              check[0] < line.length &&
              check[1] < input.length
          )
          .every((check) => number < input[check[1]][check[0]])
          ? number
          : []
      )
    )
    .reduce((result, current) => result + current + 1, 0);
});

type BasinMap = Array<Array<number | undefined | null>>;

export const task2 = withLines((data) => {
  const input = data.map((line) =>
    Array.from(line).map((char) => char !== "9")
  );

  const map: BasinMap = input.map((line) =>
    line.map((item) => (item ? undefined : null))
  );

  const groups = [];

  for (let start; (start = findNext(map)); ) {
    const neighbours = findNeighbours(map, start, []);
    neighbours.forEach(
      (neighbour) => (map[neighbour[1]][neighbour[0]] = groups.length)
    );
    groups.push(neighbours);
  }

  return groups
    .map((group) => group.length)
    .sort((a, b) => b - a)
    .slice(0, 3)
    .reduce((result, current) => result * current, 1);
});

function findNext(map: BasinMap): [number, number] | undefined {
  for (const [y, row] of map.entries()) {
    for (const [x, entry] of row.entries()) {
      if (entry === undefined) {
        return [x, y];
      }
    }
  }
}

const checks = [
  [-1, 0],
  [0, -1],
  [1, 0],
  [0, 1],
];

function findNeighbours(
  map: BasinMap,
  current: [number, number],
  previous: Array<[number, number]>
): Array<[number, number]> {
  const [x, y] = current;
  const neighbours = checks
    .map((check) => [x + check[0], y + check[1]])
    .filter(
      ([newX, newY]) =>
        newX >= 0 && newY >= 0 && newX < map[y].length && newY < map.length
    )
    .map<[number, number] | undefined>(([newX, newY]) =>
      map[newY][newX] === undefined ? [newX, newY] : undefined
    )
    .filter(isDefined)
    .filter(
      (neighbour) =>
        !previous.some(
          (prev) => prev[0] === neighbour[0] && prev[1] === neighbour[1]
        )
    );

  previous.push(current, ...neighbours);

  return [
    current,
    ...neighbours.flatMap((neighbour) =>
      findNeighbours(map, neighbour, previous)
    ),
  ];
}
