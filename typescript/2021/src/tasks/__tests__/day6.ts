import { realDataPath, test, testDataPath } from "../../testUtil";
import { task1, task2 } from "../day6";

describe("task 1", () => {
  it("test data", test(task1, testDataPath("day6"), 5934));
  it("real data", test(task1, realDataPath("day6"), 355386));
});

describe("task 2", () => {
  it("test data", test(task2, testDataPath("day6"), 26984457539));
  it("real data", test(task2, realDataPath("day6"), 1613415325809));
});
