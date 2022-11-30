import { withLines } from "../util/data";

const enum commands {
  FORWARD = "forward",
  UP = "up",
  DOWN = "down",
}

export const task1 = withLines((data) => {
  const result = data.reduce(
    (result, current) => {
      const parts = current.split(" "),
        command = parts[0],
        amount = parseInt(parts[1], 10);

      switch (command) {
        case commands.FORWARD:
          return {
            ...result,
            horizontal: result.horizontal + amount,
          };
        case commands.UP:
          return {
            ...result,
            depth: result.depth - amount,
          };
        case commands.DOWN:
          return {
            ...result,
            depth: result.depth + amount,
          };
        default:
          throw Error("Unknown command");
      }
    },
    {
      horizontal: 0,
      depth: 0,
    }
  );

  return result.horizontal * result.depth;
});

export const task2 = withLines((data) => {
  const result = data.reduce(
    (result, current) => {
      const parts = current.split(" "),
        command = parts[0],
        amount = parseInt(parts[1], 10);

      switch (command) {
        case commands.FORWARD:
          return {
            ...result,
            horizontal: result.horizontal + amount,
            depth: result.depth + result.aim * amount,
          };
        case commands.UP:
          return {
            ...result,
            aim: result.aim - amount,
          };
        case commands.DOWN:
          return {
            ...result,
            aim: result.aim + amount,
          };
        default:
          throw Error("Unknown command");
      }
    },
    {
      horizontal: 0,
      depth: 0,
      aim: 0,
    }
  );

  return result.horizontal * result.depth;
});
