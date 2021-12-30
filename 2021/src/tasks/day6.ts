import { range } from "../util";
import { withNumbers } from "../parsers";

export const task1 = withNumbers((data) => {
  return solve(data, 80);
}, ",");

export const task2 = withNumbers((data) => {
  return solve(data, 256);
}, ",");

function solve(fish: number[], iterations: number) {
  const buckets = Array.from(
    { length: 9 },
    (_, index) => fish.filter((item) => item === index).length
  );

  range(1, iterations).forEach(() => {
    const newFish = buckets.shift()!;
    buckets.push.apply(buckets, [newFish]);
    buckets[6] += newFish;
  });

  return buckets.reduce((result, current) => result + current);
}
