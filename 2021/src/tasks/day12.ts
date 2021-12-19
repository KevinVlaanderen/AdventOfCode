import { withLines } from "../parsers";

const enum Size {
  SMALL,
  LARGE,
}

type Cave = { size: Size; name: any; connections: string[] };

export const task1 = withLines((data) => {
  const connections = data.map((line) => line.split("-"));

  const caves = Array.from(
    new Set(connections.flatMap((connection) => connection))
  ).reduce(
    (result, cave) => ({
      ...result,
      [cave]: {
        size: cave === cave.toLowerCase() ? Size.SMALL : Size.LARGE,
        connections: connections
          .filter((connection) => connection.some((x) => cave === x))
          .map((connection) =>
            connection[0] === cave ? connection[1] : connection[0]
          ),
      },
    }),
    {}
  );

  return findRoutes("start", "end", caves).length;
});

function findRoutes(
  from: string,
  to: string,
  caves: { [p: string]: Cave },
  route: string[] = []
): Array<string[]> {
  return from === to
    ? [[...route, to]]
    : caves[from].connections
        .filter(
          (connection) =>
            caves[connection].size === Size.LARGE || !route.includes(connection)
        )
        .flatMap((connection) =>
          findRoutes(connection, to, caves, [...route, from])
        );
}
