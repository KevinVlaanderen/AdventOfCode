import { task1, task2 } from "../day4";
import * as path from "path";
import { loadData } from "../../util";

describe("task1", () => {
  it("gives the right answer for test data a", () => {
    const data = loadData(path.resolve(__dirname, "data/day4a"));
    expect(task1(data)).toBe(4512);
  });

  it("gives the right answer for the input", () => {
    const data = loadData("data/day4");
    expect(task1(data)).toBe(8580);
  });
});

describe("task2", () => {
  it("gives the right answer for test data a", () => {
    const data = loadData(path.resolve(__dirname, "data/day4a"));
    expect(task2(data)).toBe(1924);
  });

  it("gives the right answer for the input", () => {
    const data = loadData("data/day4");
    expect(task2(data)).toBe(9576);
  });
});
