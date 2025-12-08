import { Task } from "./types";

export function withNumbers<T>(
  task: (data: number[]) => T,
  separatedBy: string = "\n"
): Task<T> {
  return (data: string) =>
    task(
      data
        .split(separatedBy)
        .filter((entry) => entry !== "")
        .map((entry) => parseInt(entry, 10))
    );
}

export function withLines<T>(task: (data: string[]) => T): Task<T> {
  return (data: string) =>
    task(data.split("\n").filter((entry) => entry !== ""));
}

export function withBlocks<T>(task: (data: string[]) => T): Task<T> {
  return (data: string) =>
    task(data.split("\n\n").map((block) => block.trim()));
}
