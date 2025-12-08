import os

from runit.loaders.csv_loader import CSVLoader
from runit.problem import Problem


class CorruptionChecksum(Problem):
    title = "Corruption Checksum"

    # Data shape
    # 5 1 9 5
    # 7 5 3
    # 2 4 6 8
    @property
    def data(self):
        path = os.path.normpath(os.path.join(os.path.dirname(__file__), './data/data.txt'))

        lines = []

        for row in CSVLoader(path).load_async():
            row_numbers = [int(x) for x in row]
            lines.append(row_numbers)

        return lines
