import { realDataPath, test, testDataPath } from "../../util/test";
import { task1, task2 } from "../day14";

describe("task 1", () => {
  it("test data", test(task1, testDataPath("day14"), 1588));
  it("real data", test(task1, realDataPath("day14"), 3009));
});

describe("task 2", () => {
  it("test data", test(task2, testDataPath("day14"), 2188189693529));
  it("real data", test(task2, realDataPath("day14"), 3459822539451));
});
