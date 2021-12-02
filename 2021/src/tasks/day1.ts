import { Task } from "../types";
import { window } from "../util";

export const task1: Task = (data) => {
  const numbers = data.map((entry) => parseInt(entry, 10));
  let count = 0;

  for (const measurements of window(numbers, 2)) {
    if (measurements[1] > measurements[0]) {
      count += 1;
    }
  }

  return count;
};

export const task2: Task = (data) => {
  const numbers = data.map((entry) => parseInt(entry, 10));
  let previous,
    count = 0;

  for (const measurements of window(numbers, 3)) {
    const sum = measurements.reduce(
      (result, measurement) => result + measurement
    );
    if (previous && sum > previous) {
      count += 1;
    }
    previous = sum;
  }

  return count;
};
