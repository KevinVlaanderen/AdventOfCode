import { withBlocks } from "../util/data";

type Card = Array<Array<{ number: number; foundAt: number }>>;
type CardOrder = {
  card: Card;
  doneAt: number;
};

export const task1 = withBlocks((data) => {
  return solve(data, (cards) => {
    const cardsInOrder = orderCards(cards);
    return cardsInOrder[0];
  });
});

export const task2 = withBlocks((data) => {
  return solve(data, (cards) => {
    const cardsInOrder = orderCards(cards);
    return cardsInOrder[cardsInOrder.length - 1];
  });
});

function solve(
  data: string[],
  pickCard: (cards: Array<Card>) => CardOrder
): number {
  const numbers = parseNumbers(data.shift()!);
  const cards: Array<Card> = parseCards(data, numbers);

  return calculateScore(pickCard(cards), numbers);
}

function parseNumbers(block: string) {
  return block.split(",").map((item) => parseInt(item, 10));
}

function parseCards(blocks: string[], numbers: number[]) {
  return blocks.map((block) =>
    block.split("\n").map((line) =>
      line.match(/\d+/g)!.map((item) => {
        const number = parseInt(item, 10);
        return { number, foundAt: numbers.indexOf(number) };
      }, {})
    )
  );
}

function orderCards(cards: Array<Card>) {
  return cards
    .map((card) => ({
      card: card,
      doneAt: Math.min(
        //rows
        ...card
          .filter((line) => line.every((item) => item.foundAt >= 0))
          .map((line) => Math.max(...line.map((item) => item.foundAt))),
        // columns
        ...Array.from(Array(5).keys())
          .filter((column) => card.every((line) => line[column].foundAt >= 0))
          .map((column) =>
            Math.max(...card.map((line) => line[column].foundAt))
          )
      ),
    }))
    .sort((a, b) => a.doneAt - b.doneAt);
}

function calculateScore(
  firstCard: { doneAt: number; card: Card },
  numbers: number[]
) {
  return (
    firstCard.card
      .flatMap((line) => line)
      .filter((item) => item.foundAt > firstCard.doneAt)
      .reduce((result, current) => result + current.number, 0) *
    numbers[firstCard.doneAt]
  );
}
