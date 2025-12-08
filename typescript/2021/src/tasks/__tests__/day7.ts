import { realDataPath, test, testDataPath } from "../../testUtil";
import { task1, task2 } from "../day7";

describe("task 1", () => {
  it("test data", test(task1, testDataPath("day7"), 37));
  it("real data", test(task1, realDataPath("day7"), 340056));
});

describe("task 2", () => {
  it("test data", test(task2, testDataPath("day7"), 168));
  it("real data", test(task2, realDataPath("day7"), 96592275));
});
