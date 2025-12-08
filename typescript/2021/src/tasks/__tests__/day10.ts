import { realDataPath, test, testDataPath } from "../../testUtil";
import { task1, task2 } from "../day10";

describe("task 1", () => {
  it("test data", test(task1, testDataPath("day10"), 26397));
  it("real data", test(task1, realDataPath("day10"), 341823));
});

describe("task 2", () => {
  it("test data", test(task2, testDataPath("day10"), 288957));
  it("real data", test(task2, realDataPath("day10"), 2801302861));
});
