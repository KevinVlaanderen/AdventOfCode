import { task1, task2 } from "../day7";
import * as path from "path";
import { loadData } from "../../util";

describe("task1", () => {
  it("gives the right answer for test data a", () => {
    const data = loadData(path.resolve(__dirname, "data/day7a"));
    expect(task1(data)).toBe(37);
  });

  it("gives the right answer for the input", () => {
    const data = loadData("data/day7");
    expect(task1(data)).toBe(340056);
  });
});

describe("task2", () => {
  it("gives the right answer for test data a", () => {
    const data = loadData(path.resolve(__dirname, "data/day7a"));
    expect(task2(data)).toBe(168);
  });

  it("gives the right answer for the input", () => {
    const data = loadData("data/day7");
    expect(task2(data)).toBe(96592275);
  });
});
