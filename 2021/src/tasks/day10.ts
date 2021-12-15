import { withLines } from "../parsers";

interface Result {
  result: string;
}

interface CorruptedResult extends Result {
  result: "corrupted";
  expected: string;
  received: string;
}

interface IncompleteResult extends Result {
  result: "incomplete";
  expected: string[];
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

const openChars = pairs.map((pair) => pair[0]);
const closeChars = pairs.map((pair) => pair[1]);

export const task1 = withLines((data) => {
  return data
    .map(checkLine)
    .filter(isCorrupted)
    .map(mapCorruptedResultToScore)
    .reduce((result, current) => result + current, 0);
});

export const task2 = withLines((data) => {
  const scores = data
    .map(checkLine)
    .filter(isIncomplete)
    .map(mapIncompleteResultToScore)
    .sort((a, b) => a - b);

  return scores[(scores.length - 1) / 2];
});

function checkLine(line: string): Result {
  const stack: Array<string> = [];

  for (const char of Array.from(line)) {
    if (openChars.includes(char)) {
      stack.push(char);
    } else if (closeChars.includes(char)) {
      const expected = closingCharFor(stack[stack.length - 1]);

      if (expected !== char) {
        return corrupted(expected, char);
      } else {
        stack.pop();
      }
    }
  }

  if (stack.length > 0) {
    return incomplete(stack);
  }

  return success();
}

function closingCharFor(openChar: string) {
  const pair = pairs.find((pair) => pair[0] === openChar);
  if (!pair) throw new Error();
  return pair[1];
}

function corrupted(expected: string, received: string): CorruptedResult {
  return { result: "corrupted", expected, received };
}

function incomplete(unclosedPairs: string[]): IncompleteResult {
  return {
    result: "incomplete",
    expected: unclosedPairs
      .reverse()
      .map((char) => pairs.find((pair) => pair[0] === char)![1]),
  };
}

function success(): SuccessResult {
  return { result: "success" };
}

function isCorrupted(result: Result): result is CorruptedResult {
  return result.result === "corrupted";
}

function isIncomplete(result: Result): result is IncompleteResult {
  return result.result === "incomplete";
}

function mapCorruptedResultToScore(result: CorruptedResult): number {
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

function mapIncompleteResultToScore(result: IncompleteResult): number {
  return result.expected.reduce((result, current) => {
    switch (current) {
      case ")":
        return result * 5 + 1;
      case "]":
        return result * 5 + 2;
      case "}":
        return result * 5 + 3;
      case ">":
        return result * 5 + 4;
      default:
        throw new Error();
    }
  }, 0);
}
