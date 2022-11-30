import { realDataPath, test, testDataPath } from "../../util/test";
import { task1, task2 } from "../day1";

describe("task 1", () => {
  it("test data", test(task1, testDataPath("day1"), 7));
  it("real data", test(task1, realDataPath("day1"), 1266));
});

describe("task 2", () => {
  it("test data", test(task2, testDataPath("day1"), 5));
  it("real data", test(task2, realDataPath("day1"), 1217));
});
