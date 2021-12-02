import { Task } from "../types";

export const task1: Task = (data) => {
  const result = data.reduce(
    (result, current) => {
      const [direction, amount] = current.split(" ");
      return {
        horizontal:
          direction === "forward"
            ? result.horizontal + parseInt(amount, 10)
            : result.horizontal,
        depth:
          direction === "up"
            ? result.depth - parseInt(amount, 10)
            : direction === "down"
            ? result.depth + parseInt(amount, 10)
            : result.depth,
      };
    },
    {
      horizontal: 0,
      depth: 0,
    }
  );

  return result.horizontal * result.depth;
};
