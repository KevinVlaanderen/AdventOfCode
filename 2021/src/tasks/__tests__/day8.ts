import { task1 } from "../day8";
import * as path from "path";
import { loadData } from "../../util";

describe("task1", () => {
  it("gives the right answer for test data a", () => {
    const data = loadData(path.resolve(__dirname, "data/day8a"));
    expect(task1(data)).toBe(26);
  });

  it("gives the right answer for the input", () => {
    const data = loadData("data/day8");
    expect(task1(data)).toBe(539);
  });
});
