import { Task } from "../types";
import { zip } from "../util";

export const task1: Task = (data) => {
  const amount = data.length;
  const sums = data
    .map((number) => Array.from(number).map((bit) => parseInt(bit, 2)))
    .reduce((previous, current) =>
      zip(previous, current).map(([a, b]) => a + b)
    );

  const mostCommonBits = sums.map((sum) => (sum > amount / 2 ? 1 : 0));
  const leastCommonBits = mostCommonBits.map((bit) => (bit ? 0 : 1));

  const gamma = parseInt(mostCommonBits.join(""), 2);
  const epsilon = parseInt(leastCommonBits.join(""), 2);

  return gamma * epsilon;
};

export const task2: Task = (data) => {
  const numbers = data.map((number) =>
    Array.from(number).map((bit) => parseInt(bit, 2))
  );

  const oxygenGeneratorRating = parseInt(filter(numbers, "most").join(""), 2);
  const co2ScrubberRating = parseInt(filter(numbers, "least").join(""), 2);

  return oxygenGeneratorRating * co2ScrubberRating;
};

function filter(
  data: Array<number[]>,
  type: "least" | "most",
  index: number = 0
): number[] {
  if (!data || index >= data[0].length) {
    throw new Error();
  }

  const amount = data.length;
  const sum = data.reduce((previous, current) => current[index] + previous, 0);
  const criteria =
    type === "least" ? (sum < amount / 2 ? 1 : 0) : sum >= amount / 2 ? 1 : 0;
  const filteredData = data.filter((item) => item[index] === criteria);

  return filteredData.length === 1
    ? filteredData[0]
    : filter(filteredData, type, index + 1);
}
