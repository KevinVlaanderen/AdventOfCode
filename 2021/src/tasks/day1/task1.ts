import { Task } from "../../types";

export const task1: Task = (data) => {
  const measurements = data.map((entry) => parseInt(entry, 10));
  let previous,
    count = 0;

  for (const measurement of measurements) {
    if (previous && measurement > previous) {
      count += 1;
    }
    previous = measurement;
  }

  return count;
};
