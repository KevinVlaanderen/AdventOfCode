import { Task } from "../types";
import { range } from "../util";

export const task1: Task = (data) => {
  return breed(
    data.split(",").map((item) => parseInt(item, 10)),
    80
  );
};

export const task2: Task = (data) => {
  return breed(
    data.split(",").map((item) => parseInt(item, 10)),
    256
  );
};

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
