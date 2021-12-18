import { loadData } from "./util";
import path from "path";

export function test(
  task: (data: string) => number,
  path: string,
  expected: number
) {
  return () => {
    const data = loadData(path);
    expect(task(data)).toBe(expected);
  };
}

export function testDataPath(name: string): string {
  return path.resolve(__dirname, "tasks/__tests__/data", name);
}

export function realDataPath(name: string): string {
  return `data/${name}`;
}
