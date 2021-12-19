import { withLines } from "../parsers";

const enum Size {
  SMALL,
  LARGE,
}

type Cave = { size: Size; name: any; connections: string[] };

export const task1 = withLines((data) => {
  const connections = getConnections(data);
  const caves = getCaves(connections);

  return findRoutes("start", "end", caves).length;
});

export const task2 = withLines((data) => {
  const connections = getConnections(data);
  const caves = getCaves(connections);

  const smallCaves = Object.entries(caves)
    .filter(
      ([name, cave]) =>
        cave.size === Size.SMALL && name !== "start" && name !== "end"
    )
    .map(([name]) => name);

  return new Set(
    smallCaves.flatMap((exception) =>
      findRoutes("start", "end", caves, exception).map((path) => path.join(","))
    )
  ).size;
});

function getConnections(data: string[]) {
  return data.map((line) => line.split("-"));
}

function getCaves(connections: string[][]): { [cave: string]: Cave } {
  return Array.from(
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
}

function findRoutes(
  from: string,
  to: string,
  caves: { [p: string]: Cave },
  exception?: string,
  route: string[] = []
): Array<string[]> {
  return from === to
    ? [[...route, to]]
    : caves[from].connections
        .filter(
          (connection) =>
            caves[connection].size === Size.LARGE ||
            (connection === exception &&
              route.filter((visited) => connection === visited).length <= 1) ||
            !route.includes(connection)
        )
        .flatMap((connection) =>
          findRoutes(connection, to, caves, exception, [...route, from])
        );
}
