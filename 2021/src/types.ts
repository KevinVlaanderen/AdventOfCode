export type Task<T> = (data: string) => T;

export interface Results<T> {
  [name: string]: T;
}
