import { Task } from "../types";
import { zip } from "../util";

export const task1: Task = (data) => {
  const amount = data.length;
  const sums = data
    .map((number) => Array.from(number).map((bit) => parseInt(bit, 10)))
    .reduce((previous, current) =>
      zip(previous, current).map(([a, b]) => a + b)
    );

  const mostCommonBits = sums.map((sum) => (sum > amount / 2 ? 1 : 0));
  const leastCommonBits = mostCommonBits.map((bit) => (bit ? 0 : 1));

  const gamma = parseInt(mostCommonBits.join(""), 2);
  const epsilon = parseInt(leastCommonBits.join(""), 2);

  return gamma * epsilon;
};
