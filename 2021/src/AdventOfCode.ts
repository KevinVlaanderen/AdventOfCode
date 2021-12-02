import { tasks } from "./tasks";
import { readFileSync } from "fs";

interface Results<T> {
  [name: string]: T;
}

const results = Object.entries(tasks).reduce(
  (dayResult: Results<Results<number>>, [dayName, tasks]) => {
    const data = readFileSync(`data/${dayName}`, "utf-8")
      .split("\n")
      .filter((entry) => entry !== "");

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
