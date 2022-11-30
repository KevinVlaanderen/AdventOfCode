import { realDataPath, test, testDataPath } from "../../util/test";
import { task1, task2 } from "../day3";

describe("task 1", () => {
  it("test data", test(task1, testDataPath("day3"), 198));
  it("real data", test(task1, realDataPath("day3"), 3847100));
});

describe("task 2", () => {
  it("test data", test(task2, testDataPath("day3"), 230));
  it("real data", test(task2, realDataPath("day3"), 4105235));
});
