import { Task } from "../types";

type Card = Array<Array<{ number: number; foundAt: number }>>;

export const task1: Task = (data) => {
  const blocks = data.split("\n\n").map((block) => block.trim());

  const numbers = blocks
    .shift()!
    .split(",")
    .map((item) => parseInt(item, 10));

  const cards: Array<Card> = blocks.map((block) =>
    block.split("\n").map((line) =>
      line.match(/\d+/g)!.map((item) => {
        const number = parseInt(item, 10);
        return { number, foundAt: numbers.indexOf(number) };
      }, {})
    )
  );

  const cardsInOrder = cards
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

  return (
    cardsInOrder[0].card
      .flatMap((line) => line)
      .filter((item) => item.foundAt > cardsInOrder[0].doneAt)
      .reduce((result, current) => result + current.number, 0) *
    numbers[cardsInOrder[0].doneAt]
  );
};
