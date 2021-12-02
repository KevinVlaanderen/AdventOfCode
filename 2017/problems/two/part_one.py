import operator
from functools import reduce
from typing import List

from runit.strategy import Strategy


class PartOne(Strategy):
    title = "Part One"

    def execute(self, numbers: List[List[int]]) -> int:
        result = reduce(operator.add, [max(row) - min(row) for row in numbers])

        yield result

