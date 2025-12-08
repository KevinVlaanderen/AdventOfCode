import { realDataPath, test, testDataPath } from "../../testUtil";
import { task1, task2 } from "../day4";

describe("task 1", () => {
  it("test data", test(task1, testDataPath("day4"), 4512));
  it("real data", test(task1, realDataPath("day4"), 8580));
});

describe("task 2", () => {
  it("test data", test(task2, testDataPath("day4"), 1924));
  it("real data", test(task2, realDataPath("day4"), 9576));
});
