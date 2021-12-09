import { Task } from "../types";
import { range } from "../util";

export const task1: Task = (data) => {
  const positions = data.split(",").map((item) => parseInt(item, 10));
  return range(0, Math.max(...positions), 1)
    .map((index) =>
      positions.reduce(
        (result, current) => result + Math.abs(current - index),
        0
      )
    )
    .sort((a, b) => a - b)[0];
};
