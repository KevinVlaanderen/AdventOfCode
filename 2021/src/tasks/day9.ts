import { withLines } from "../parsers";

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
