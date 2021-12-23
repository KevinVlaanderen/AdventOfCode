import { realDataPath, test, testDataPath } from "../../testUtil";
import { task1, task2 } from "../day13";

describe("task 1", () => {
  it("test data", test(task1, testDataPath("day13"), 17));
  it("real data", test(task1, realDataPath("day13"), 687));
});

describe("task 2", () => {
  it("real data", test(task2, realDataPath("day13"), "FGKCKBZG"));
});
