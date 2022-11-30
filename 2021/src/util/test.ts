import path from "path";
import { loadData } from "./data";

export function test<T>(task: (data: string) => T, path: string, expected: T) {
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
