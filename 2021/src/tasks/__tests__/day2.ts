import { task1 } from "../day2";
import * as path from "path";
import { loadData } from "../../util";

describe("task1", () => {
  it("gives the right answer for test data a", () => {
    const data = loadData(path.resolve(__dirname, "data/day2/task1a"));
    expect(task1(data)).toBe(150);
  });

  it("gives the right answer for the input", () => {
    const data = loadData("data/day2");
    expect(task1(data)).toBe(1524750);
  });
});
