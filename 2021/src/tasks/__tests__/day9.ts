import { task1 } from "../day9";
import * as path from "path";
import { loadData } from "../../util";

describe("task1", () => {
  it("gives the right answer for test data a", () => {
    const data = loadData(path.resolve(__dirname, "data/day9a"));
    expect(task1(data)).toBe(15);
  });

  it("gives the right answer for the input", () => {
    const data = loadData("data/day9");
    expect(task1(data)).toBe(452);
  });
});
