import { realDataPath, test, testDataPath } from "../../util/test";
import { task1, task2 } from "../day8";

describe("task 1", () => {
  it("test data", test(task1, testDataPath("day8"), 26));
  it("real data", test(task1, realDataPath("day8"), 539));
});

describe("task 2", () => {
  it("test data", test(task2, testDataPath("day8"), 61229));
  it("real data", test(task2, realDataPath("day8"), 1084606));
});
