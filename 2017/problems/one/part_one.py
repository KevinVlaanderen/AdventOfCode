import operator
from functools import reduce

from runit.strategy import Strategy


class PartOne(Strategy):
    title = "Part One"

    def execute(self, numbers: str) -> int:
        shifted = numbers[len(numbers)-1] + numbers[:-1]

        result = reduce(operator.add, [int(x) for x, y in zip(numbers, shifted) if x == y], initialized=0)

        yield result

