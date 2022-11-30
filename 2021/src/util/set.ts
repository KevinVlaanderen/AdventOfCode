export function subtract<T>(a: Set<T>, ...b: Set<T>[]): Set<T> {
  return new Set([...a].filter((x) => !b.some((y) => y.has(x))));
}

export function equal<T>(a: Set<T>, b: Set<T>): boolean {
  return a.size === b.size && [...a].every((v) => b.has(v));
}
