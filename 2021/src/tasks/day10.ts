import { withLines } from "../parsers";

interface Result {
  result: string;
}

interface CorruptedResult extends Result {
  result: "corrupted";
  reason: string;
  expected: string;
  received: string;
}

interface IncompleteResult extends Result {
  result: "incomplete";
  reason: string;
}

interface SuccessResult extends Result {
  result: "success";
}

const pairs = [
  ["(", ")"],
  ["[", "]"],
  ["{", "}"],
  ["<", ">"],
];
const openingChars = pairs.map((pair) => pair[0]);
const closingChars = pairs.map((pair) => pair[1]);

export const task1 = withLines((data) => {
  return data
    .map((line) => {
      const stack: Array<string> = [];

      for (const char of Array.from(line)) {
        if (openingChars.includes(char)) {
          stack.push(char);
        } else {
          const expectedPair = pairs.find(
            (pair) => pair[0] === stack[stack.length - 1]
          )!;

          if (!closingChars.includes(char)) {
            return corrupted("unknown character", expectedPair[1], char);
          } else if (expectedPair[1] !== char) {
            return corrupted("unexpected character", expectedPair[1], char);
          } else {
            stack.pop();
          }
        }
      }

      if (stack.length > 0) {
        return incomplete(stack);
      }

      return success();
    })
    .filter(isCorrupted)
    .map(mapResultToScore)
    .reduce((result, current) => result + current, 0);
});

function corrupted(
  reason: string,
  expected: string,
  received: string
): CorruptedResult {
  return { result: "corrupted", reason, expected, received };
}

function incomplete(unclosedPairs: string[]): IncompleteResult {
  return {
    result: "incomplete",
    reason: `unclosed pairs: ${unclosedPairs.join()}`,
  };
}

function success(): SuccessResult {
  return { result: "success" };
}

function isCorrupted(result: Result): result is CorruptedResult {
  return result.result === "corrupted";
}

function mapResultToScore(result: CorruptedResult) {
  switch (result.received) {
    case ")":
      return 3;
    case "]":
      return 57;
    case "}":
      return 1197;
    case ">":
      return 25137;
    default:
      throw new Error();
  }
}
