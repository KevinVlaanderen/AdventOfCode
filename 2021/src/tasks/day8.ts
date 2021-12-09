import { withLines } from "../parsers";

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
