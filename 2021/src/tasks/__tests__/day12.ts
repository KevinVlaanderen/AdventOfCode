import { realDataPath, test, testDataPath } from "../../testUtil";
import { task1, task2 } from "../day12";

describe("task 1", () => {
  it("test data a", test(task1, testDataPath("day12a"), 10));
  it("test data b", test(task1, testDataPath("day12b"), 19));
  it("test data c", test(task1, testDataPath("day12c"), 226));
  it("real data", test(task1, realDataPath("day12"), 3708));
});

describe("task 2", () => {
  it("test data a", test(task2, testDataPath("day12a"), 36));
  it("test data b", test(task2, testDataPath("day12b"), 103));
  it("test data c", test(task2, testDataPath("day12c"), 3509));
  it("real data", test(task2, realDataPath("day12"), 93858));
});
