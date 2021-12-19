import { realDataPath, test, testDataPath } from "../../testUtil";
import { task1 } from "../day11";

describe("task 1", () => {
  it("test data", test(task1, testDataPath("day11"), 1656));
  it("real data", test(task1, realDataPath("day11"), 1644));
});
