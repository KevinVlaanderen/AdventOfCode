import { realDataPath, test, testDataPath } from "../../util/test";
import { task1, task2 } from "../day9";

describe("task 1", () => {
  it("test data", test(task1, testDataPath("day9"), 15));
  it("real data", test(task1, realDataPath("day9"), 452));
});

describe("task 2", () => {
  it("test data", test(task2, testDataPath("day9"), 1134));
  it("real data", test(task2, realDataPath("day9"), 1263735));
});
