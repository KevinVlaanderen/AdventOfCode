import math
from runit.strategy import Strategy


class PartOne(Strategy):
    title = "Part One"

    def execute(self, start: int) -> int:
        block_size = math.ceil(start**0.5)
        if block_size % 2 == 0:
            block_size += 1
        biggest_in_current_layer = block_size ** 2
        biggest_in_previous_layer = (block_size - 2) ** 2
        difference_between_layers = biggest_in_current_layer - biggest_in_previous_layer
        side_length = difference_between_layers / 4
        progress_within_layer = start - biggest_in_previous_layer
        distance_from_center = abs(progress_within_layer % side_length - (side_length / 2))

        yield int(block_size / 2) + distance_from_center
