import { withLines } from "../parsers";
import { equal, subtract } from "../util";

const segmentsPerNumber = {
  0: 6,
  1: 2,
  2: 5,
  3: 5,
  4: 4,
  5: 5,
  6: 6,
  7: 3,
  8: 7,
  9: 6,
};

export const task1 = withLines((data) => {
  return data
    .flatMap((line) => line.match(/[abcdefg]+/g)!.slice(-4))
    .filter(
      (digit) =>
        digit.length === segmentsPerNumber[1] ||
        digit.length === segmentsPerNumber[4] ||
        digit.length === segmentsPerNumber[7] ||
        digit.length === segmentsPerNumber[8]
    ).length;
});

export const task2 = withLines((data) => {
  const input = data.map((line) => {
    const parts = line
      .split("|")
      .map((part) =>
        part
          .match(/[abcdefg]+/g)!
          .map((pattern) => new Set(Array.from(pattern)))
      );
    return {
      signal: [...new Set(parts[0])],
      output: parts[1],
    };
  });

  return input
    .map((entry) => {
      const numberLookup: {
        [number: number]: Set<string>;
      } = {};

      numberLookup[1] = entry.signal.find((word) => word.size === 2)!;
      numberLookup[4] = entry.signal.find((word) => word.size === 4)!;
      numberLookup[7] = entry.signal.find((word) => word.size === 3)!;
      numberLookup[8] = entry.signal.find((word) => word.size === 7)!;

      const words069 = entry.signal.filter((word) => word.size === 6);
      const words235 = entry.signal.filter((word) => word.size === 5);

      numberLookup[0] = words069.find(
        (word) =>
          subtract(word, numberLookup[4]).size === 3 &&
          subtract(word, numberLookup[7]).size === 3
      )!;
      numberLookup[2] = words235.find(
        (word) => subtract(word, numberLookup[4]).size === 3
      )!;
      numberLookup[3] = words235.find(
        (word) =>
          subtract(word, numberLookup[1]).size === 3 &&
          subtract(word, numberLookup[4]).size === 2
      )!;

      numberLookup[5] = words235.find(
        (word) =>
          subtract(word, numberLookup[1]).size === 4 &&
          subtract(word, numberLookup[4]).size === 2
      )!;
      numberLookup[6] = words069.find(
        (word) =>
          subtract(word, numberLookup[4]).size === 3 &&
          subtract(word, numberLookup[7]).size === 4
      )!;
      numberLookup[9] = words069.find(
        (word) => subtract(word, numberLookup[4]).size === 2
      )!;

      return parseInt(
        entry.output
          .map(
            (word) =>
              Object.entries(numberLookup).find((x) => equal(x[1], word))![0]
          )
          .join(""),
        10
      );
    })
    .reduce((result, current) => result + current);
});
