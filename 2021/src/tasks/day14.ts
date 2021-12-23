import { withBlocks } from "../parsers";
import { range, window } from "../util";

export const task1 = withBlocks((data) => {
  const rules = parseRules(data[1]);

  const result = range(1, 10).reduce((template, _) => {
    const pairs = Array.from(window(template, 2));

    return pairs.flatMap((pair, index) => {
      const newElement = rules[pair.join("")];
      return index < pairs.length - 1
        ? [pair[0], newElement]
        : [pair[0], newElement, pair[1]];
    });
  }, data[0].split(""));

  const uniqueElements = new Set(result);
  const counts = Array.from(uniqueElements)
    .map<[string, number]>((element) => [
      element,
      result.filter((item) => item === element).length,
    ])
    .sort((a, b) => b[1] - a[1]);

  return counts[0][1] - counts[counts.length - 1][1];
});

function parseRules(data: string): { [key: string]: string } {
  return data.split("\n").reduce((result, line) => {
    const matches = line.match(/(..) -> (.)/)!;
    return {
      ...result,
      [matches[1]]: matches[2],
    };
  }, {});
}
