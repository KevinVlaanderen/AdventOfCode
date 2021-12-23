import { realDataPath, test, testDataPath } from "../../testUtil";
import { task1 } from "../day14";

describe("task 1", () => {
  it("test data", test(task1, testDataPath("day14"), 1588));
  it("real data", test(task1, realDataPath("day14"), 3009));
});
