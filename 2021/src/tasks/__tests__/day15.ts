import { realDataPath, test, testDataPath } from "../../util/test";
import { task1 } from "../day15";

describe("task 1", () => {
  it("test data", test(task1, testDataPath("day15"), 40));
  it("real data", test(task1, realDataPath("day15"), undefined));
});
