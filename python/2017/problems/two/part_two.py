import operator
from functools import reduce
from typing import List

from runit.strategy import Strategy


class PartTwo(Strategy):
    title = "Part Two"

    def execute(self, numbers: List[List[int]]) -> int:
        result = reduce(operator.add, [high / low
                                       for row in numbers
                                       for high in sorted(row, reverse=True)
                                       for low in sorted(row)
                                       if high > low and high % low == 0])

        yield result

