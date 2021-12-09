import { task1, task2 } from "../day6";
import * as path from "path";
import { loadData } from "../../util";

describe("task1", () => {
  it("gives the right answer for test data a", () => {
    const data = loadData(path.resolve(__dirname, "data/day6/task1a"));
    expect(task1(data)).toBe(5934);
  });

  it("gives the right answer for the input", () => {
    const data = loadData("data/day6");
    expect(task1(data)).toBe(355386);
  });
});

describe("task2", () => {
  it("gives the right answer for test data a", () => {
    const data = loadData(path.resolve(__dirname, "data/day6/task2a"));
    expect(task2(data)).toBe(26984457539);
  });

  it("gives the right answer for the input", () => {
    const data = loadData("data/day6");
    expect(task2(data)).toBe(1613415325809);
  });
});
