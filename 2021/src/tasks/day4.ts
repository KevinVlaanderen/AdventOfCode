import { Task } from "../types";

type Card = Array<Array<{ number: number; foundAt: number }>>;

export const task1: Task = (data) => {
  const blocks = data.split("\n\n").map((block) => block.trim());

  const numbers = parseNumbers(blocks.shift()!);
  const cards: Array<Card> = parseCards(blocks, numbers);
  const cardsInOrder = orderCards(cards);

  const firstCard = cardsInOrder[0];

  return calculateScore(firstCard, numbers);
};

export const task2: Task = (data) => {
  const blocks = data.split("\n\n").map((block) => block.trim());

  const numbers = parseNumbers(blocks.shift()!);
  const cards: Array<Card> = parseCards(blocks, numbers);
  const cardsInOrder = orderCards(cards);

  const lastCard = cardsInOrder[cardsInOrder.length - 1];

  return calculateScore(lastCard, numbers);
};

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
