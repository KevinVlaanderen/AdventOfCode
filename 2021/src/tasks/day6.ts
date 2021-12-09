import { range } from "../util";
import { withNumbers } from "../parsers";

export const task1 = withNumbers((data) => {
  return breed(data, 80);
}, ",");

export const task2 = withNumbers((data) => {
  return breed(data, 256);
}, ",");

function breed(fish: number[], iterations: number) {
  const buckets = Array.from(
    { length: 9 },
    (_, index) => fish.filter((item) => item === index).length
  );

  range(1, iterations, 1).forEach(() => {
    const newFish = buckets.shift()!;
    buckets.push.apply(buckets, [newFish]);
    buckets[6] += newFish;
  });

  return buckets.reduce((result, current) => result + current);
}
