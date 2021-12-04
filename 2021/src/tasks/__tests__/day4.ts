import { task1 } from "../day4";
import * as path from "path";
import { loadData } from "../../util";

describe("task1", () => {
  it("gives the right answer for test data a", () => {
    const data = loadData(path.resolve(__dirname, "data/day4/task1a"));
    expect(task1(data)).toBe(4512);
  });

  it("gives the right answer for the input", () => {
    const data = loadData("data/day4");
    expect(task1(data)).toBe(8580);
  });
});
