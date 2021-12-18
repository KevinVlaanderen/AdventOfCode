import { loadData } from "./util";
import * as day1 from "./tasks/day1";
import * as day2 from "./tasks/day2";
import * as day3 from "./tasks/day3";
import * as day4 from "./tasks/day4";
import * as day5 from "./tasks/day5";
import * as day6 from "./tasks/day6";
import * as day7 from "./tasks/day7";
import * as day8 from "./tasks/day8";
import * as day9 from "./tasks/day9";
import * as day10 from "./tasks/day10";

import { Results } from "./types";

const tasks = {
  day1,
  day2,
  day3,
  day4,
  day5,
  day6,
  day7,
  day8,
  day9,
  day10,
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