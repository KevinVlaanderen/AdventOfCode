import { task1, task2 } from "../day5";
import * as path from "path";
import { loadData } from "../../util";

describe("task1", () => {
  it("gives the right answer for test data a", () => {
    const data = loadData(path.resolve(__dirname, "data/day5/task1a"));
    expect(task1(data)).toBe(5);
  });

  it("gives the right answer for the input", () => {
    const data = loadData("data/day5");
    expect(task1(data)).toBe(6189);
  });
});

describe("task2", () => {
  it("gives the right answer for test data a", () => {
    const data = loadData(path.resolve(__dirname, "data/day5/task2a"));
    expect(task2(data)).toBe(12);
  });

  it("gives the right answer for the input", () => {
    const data = loadData("data/day5");
    expect(task2(data)).toBe(19164);
  });
});
