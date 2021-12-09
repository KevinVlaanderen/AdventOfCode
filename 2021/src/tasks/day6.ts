import { Task } from "../types";
import { range } from "../util";

export const task1: Task = (data) => {
  const fish = data.split(",").map((item) => parseInt(item, 10));

  range(1, 80, 1).forEach(() => {
    const newFish = fish.filter((item) => item === 0).map(() => 8);
    fish.forEach((item, index) => (fish[index] = item === 0 ? 6 : item - 1));
    fish.push.apply(fish, newFish);
  });

  return fish.length;
};
