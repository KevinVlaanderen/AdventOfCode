import { task1 } from "../day7";
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
