import { range } from "../util";
import { withNumbers } from "../parsers";

export const task1 = withNumbers((data) => {
  return calculateCosts(data, (distance) => distance)[0];
}, ",");

export const task2 = withNumbers((data) => {
  return calculateCosts(data, (distance) => (distance * (distance + 1)) / 2)[0];
}, ",");

function calculateCosts(
  positions: number[],
  costFn: (distance: number) => number
) {
  return range(0, Math.max(...positions))
    .map((index) =>
      positions.reduce(
        (result, current) => result + costFn(Math.abs(current - index)),
        0
      )
    )
    .sort((a, b) => a - b);
}
