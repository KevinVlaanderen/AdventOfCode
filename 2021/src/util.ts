import { readFileSync } from "fs";

export function loadData(path: string) {
  return readFileSync(path, "utf-8");
}

export function* window<T>(array: T[], size: number): Generator<T[]> {
  for (let index = 0; index <= array.length - size; index++) {
    yield array.slice(index, index + size);
  }
}

export function zip(arr: any[], ...arrs: Array<any[]>) {
  return arr.map((val, i) => arrs.reduce((a, arr) => [...a, arr[i]], [val]));
}

export function range(start: number, stop: number, step: number) {
  return Array.from(
    { length: Math.abs(stop - start) / step + 1 },
    (_, i) => start + i * step * (start > stop ? -1 : 1)
  );
}
