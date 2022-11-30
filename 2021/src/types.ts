export type Task<T> = (data: string) => T;

export interface Results<T> {
  [name: string]: T;
}

export function isDefined<T>(val: T | undefined | null): val is T {
  return val !== undefined && val !== null;
}
