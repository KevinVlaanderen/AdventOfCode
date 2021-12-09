import { Task } from "../types";
import { range } from "../util";

export const task1: Task = (data) => {
  const positions = data.split(",").map((item) => parseInt(item, 10));

  return calculateCosts(positions, (distance) => distance)[0];
};

export const task2: Task = (data) => {
  const positions = data.split(",").map((item) => parseInt(item, 10));

  return calculateCosts(
    positions,
    (distance) => (distance * (distance + 1)) / 2
  )[0];
};

function calculateCosts(
  positions: number[],
  costFn: (distance: number) => number
) {
  return range(0, Math.max(...positions), 1)
    .map((index) =>
      positions.reduce(
        (result, current) => result + costFn(Math.abs(current - index)),
        0
      )
    )
    .sort((a, b) => a - b);
}
