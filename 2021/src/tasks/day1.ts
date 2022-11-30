import { withNumbers } from "../util/data";
import { window } from "../util/array";

export const task1 = withNumbers((data) => {
  let count = 0;

  for (const measurements of window(data, 2)) {
    if (measurements[1] > measurements[0]) {
      count += 1;
    }
  }

  return count;
});

export const task2 = withNumbers((data) => {
  let previous,
    count = 0;

  for (const measurements of window(data, 3)) {
    const sum = measurements.reduce(
      (result, measurement) => result + measurement
    );
    if (previous && sum > previous) {
      count += 1;
    }
    previous = sum;
  }

  return count;
});
