import { loadData } from "./util";
import * as day1 from "./tasks/day1";
import { Results } from "./types";

const tasks = {
  day1,
};

const results = Object.entries(tasks).reduce(
  (dayResult: Results<Results<number>>, [dayName, tasks]) => {
    const data = loadData("data/${dayName}");

    return {
      ...dayResult,
      [dayName]: Object.entries(tasks).reduce(
        (dayResult: Results<number>, [taskName, task]) => ({
          ...dayResult,
          [taskName]: task(data),
        }),
        {}
      ),
    };
  },
  {}
);

console.log(JSON.stringify(results, null, "\t"));
