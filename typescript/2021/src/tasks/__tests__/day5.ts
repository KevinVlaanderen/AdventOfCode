import { realDataPath, test, testDataPath } from "../../testUtil";
import { task1, task2 } from "../day5";

describe("task 1", () => {
  it("test data", test(task1, testDataPath("day5"), 5));
  it("real data", test(task1, realDataPath("day5"), 6189));
});

describe("task 2", () => {
  it("test data", test(task2, testDataPath("day5"), 12));
  it("real data", test(task2, realDataPath("day5"), 19164));
});
