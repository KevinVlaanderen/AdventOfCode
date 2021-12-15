import { task1 } from "../day10";
import * as path from "path";
import { loadData } from "../../util";

describe("task1", () => {
  it("gives the right answer for test data a", () => {
    const data = loadData(path.resolve(__dirname, "data/day10a"));
    expect(task1(data)).toBe(26397);
  });

  it("gives the right answer for the input", () => {
    const data = loadData("data/day10");
    expect(task1(data)).toBe(341823);
  });
});
