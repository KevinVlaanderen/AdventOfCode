import { Task } from "./types";

export function withNumbers(
  task: (data: number[]) => number,
  separatedBy: string = "\n"
): Task {
  return (data: string) =>
    task(
      data
        .split(separatedBy)
        .filter((entry) => entry !== "")
        .map((entry) => parseInt(entry, 10))
    );
}

export function withLines(task: (data: string[]) => number): Task {
  return (data: string) =>
    task(data.split("\n").filter((entry) => entry !== ""));
}

export function withBlocks(task: (data: string[]) => number): Task {
  return (data: string) =>
    task(data.split("\n\n").map((block) => block.trim()));
}
