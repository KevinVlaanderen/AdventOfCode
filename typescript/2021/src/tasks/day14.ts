import { withBlocks } from "../parsers";
import { window } from "../util";

type Rules = {
  between: { [key: string]: string };
  pairs: { [key: string]: [string, string] };
};

export const task1 = withBlocks((data) => {
  return solve(data, 10);
});

export const task2 = withBlocks((data) => {
  return solve(data, 40);
});

function solve(data: string[], steps: number): number {
  const template = data[0];
  const rules = parseRules(data[1]);

  const counter = new Counter(...template);
  const cache = new Cache();

  counter.merge(grow(template, rules, steps, cache));

  return calculateScore(counter);
}

function parseRules(data: string): Rules {
  const rules: Rules = { between: {}, pairs: {} };
  data.split("\n").forEach((line) => {
    const matches = line.match(/(..) -> (.)/)!;
    rules.between[matches[1]] = matches[2];
    rules.pairs[matches[1]] = [
      matches[1][0] + matches[2],
      matches[2] + matches[1][1],
    ];
  });
  return rules;
}

function grow(
  template: string,
  rules: Rules,
  step: number,
  cache: Cache
): Counter {
  if (step === 0) return new Counter();

  if (cache) {
    const cachedCounter = cache.get(template, step);
    if (cachedCounter) return cachedCounter;
  }

  const counter = new Counter();

  const letters = Array.from(template);
  for (const [start, end] of window(letters, 2)) {
    if (rules.pairs[start + end]) {
      const between = rules.between[start + end];
      counter.add(between);
      counter.merge(grow(start + between + end, rules, step - 1, cache));
    }
  }

  cache.set(template, step, counter);

  return counter;
}

function calculateScore(counter: Counter) {
  const sorted = Object.values(counter.all()).sort((a, b) => b - a);
  return sorted[0] - sorted[sorted.length - 1];
}

class Counter {
  private counts: { [character: string]: number } = {};

  constructor(...character: string[]) {
    character.forEach((character) => this.add(character));
  }

  add(key: string): void {
    this.counts[key] = this.counts[key] ? this.counts[key] + 1 : 1;
  }

  get(key: string): number {
    return this.counts[key] ?? 0;
  }

  all(): { [character: string]: number } {
    return this.counts;
  }

  merge(counter: Counter): void {
    Object.entries(counter.all()).forEach(([character, n]) => {
      this.counts[character] = this.get(character) + n;
    });
  }
}

class Cache {
  private cache: { [template: string]: { [step: number]: Counter } } = {};

  set(template: string, step: number, counter: Counter): void {
    if (this.cache[template] === undefined) {
      this.cache[template] = {};
    }
    if (this.cache[template][step] === undefined) {
      this.cache[template][step] = counter;
    }
  }

  get(template: string, step: number): Counter | undefined {
    return template in this.cache ? this.cache[template][step] : undefined;
  }
}
