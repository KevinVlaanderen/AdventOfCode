export interface Runnable {
  run: () => void;
}

export type Task = (data: string[]) => number;
