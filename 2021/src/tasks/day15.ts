import { withLines } from "../util/data";
import { priorityQueue } from "../util/priorityQueue";

export const task1 = withLines((data) => {
  const tiles = data.map((line) =>
    Array.from(line).map((item) => parseInt(item, 10))
  );

  const frontier = priorityQueue<Tile>([[from, 0]]);
  const cameFrom = new TileMap<Tile | undefined>([[from, undefined]]);
  const costSoFar = new TileMap<number>([[from, 0]]);

  while (!frontier.isEmpty()) {
    const current = frontier.pop()!;

    if (current === to) {
      break;
    }

    this.tileProvider.tiles(neighbours(current)).forEach((next) => {
      const currentCost = costSoFar.get(current)!;
      const nextCost = this.getMovementCost(
        next,
        next.coordinates.q === to.coordinates.q &&
          next.coordinates.r === to.coordinates.r
      );
      const newCost = currentCost + nextCost;

      if (!costSoFar.has(next) || newCost < costSoFar.get(next)!) {
        const priority = newCost + this.heuristic(next, to);
        frontier.insert(next, priority);
        costSoFar.set(next, newCost);
        cameFrom.set(next, current);
      }
    });
  }

  return -1;
});
