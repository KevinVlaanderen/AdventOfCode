import { withNumbers } from "../util/data";
import { range } from "../util/array";

export const task1 = withNumbers((data) => {
  return solve(data, (distance) => distance)[0];
}, ",");

export const task2 = withNumbers((data) => {
  return solve(data, (distance) => (distance * (distance + 1)) / 2)[0];
}, ",");

function solve(positions: number[], costFn: (distance: number) => number) {
  return range(0, Math.max(...positions))
    .map((index) =>
      positions.reduce(
        (result, current) => result + costFn(Math.abs(current - index)),
        0
      )
    )
    .sort((a, b) => a - b);
}
