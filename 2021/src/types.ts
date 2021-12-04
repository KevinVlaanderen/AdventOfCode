export type Task = (data: string) => number;

export interface Results<T> {
  [name: string]: T;
}
