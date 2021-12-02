import { task1, task2 } from "../day1";
import * as path from "path";
import { loadData } from "../../util";

describe("task1", () => {
  it("gives the right answer for test data a", () => {
    const data = loadData(path.resolve(__dirname, "data/day1/task1a"));
    expect(task1(data)).toBe(7);
  });

  it("gives the right answer for the input", () => {
    const data = loadData("data/day1");
    expect(task1(data)).toBe(1266);
  });
});

describe("task2", () => {
  it("gives the right answer for test data a", () => {
    const data = loadData(path.resolve(__dirname, "data/day1/task2a"));
    expect(task2(data)).toBe(5);
  });

  it("gives the right answer for the input", () => {
    const data = loadData("data/day1");
    expect(task2(data)).toBe(1217);
  });
});
