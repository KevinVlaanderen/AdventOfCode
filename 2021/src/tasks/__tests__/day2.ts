import { realDataPath, test, testDataPath } from "../../testUtil";
import { task1, task2 } from "../day2";

describe("task 1", () => {
  it("test data", test(task1, testDataPath("day2"), 150));
  it("real data", test(task1, realDataPath("day2"), 1524750));
});

describe("task 2", () => {
  it("test data", test(task2, testDataPath("day2"), 900));
  it("real data", test(task2, realDataPath("day2"), 1592426537));
});
