import operator
from functools import reduce
from runit.strategy import Strategy


class PartTwo(Strategy):
    title = "Part Two"

    def execute(self, numbers: str) -> int:
        assert len(numbers) % 2 == 0

        length = len(numbers)
        firstpart, secondpart = numbers[:length // 2], numbers[length // 2:]
        shifted = secondpart + firstpart

        result = reduce(operator.add, [int(x) for x, y in zip(numbers, shifted) if x == y])

        yield result

