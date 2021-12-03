import { task1, task2 } from "../day3";
import * as path from "path";
import { loadData } from "../../util";

describe("task1", () => {
  it("gives the right answer for test data a", () => {
    const data = loadData(path.resolve(__dirname, "data/day3/task1a"));
    expect(task1(data)).toBe(198);
  });

  it("gives the right answer for the input", () => {
    const data = loadData("data/day3");
    expect(task1(data)).toBe(3847100);
  });
});

describe("task2", () => {
  it("gives the right answer for test data a", () => {
    const data = loadData(path.resolve(__dirname, "data/day3/task2a"));
    expect(task2(data)).toBe(230);
  });

  it("gives the right answer for the input", () => {
    const data = loadData("data/day3");
    expect(task2(data)).toBe(4105235);
  });
});
