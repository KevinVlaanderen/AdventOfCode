import os

from runit.problem import Problem


class InverseCaptcha(Problem):
    title = "Inverse Captcha"

    # Data shape
    # 123456789112233445566778899
    @property
    def data(self):
        path = os.path.normpath(os.path.join(os.path.dirname(__file__), './data/data.txt'))

        with open(path) as file:
            return file.readline()


