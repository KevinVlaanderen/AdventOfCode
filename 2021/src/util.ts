import { readFileSync } from "fs";

export function loadData(path: string) {
  return readFileSync(path, "utf-8")
    .split("\n")
    .filter((entry) => entry !== "");
}
