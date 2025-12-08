from problems.one.part_one import PartOne


def test_1122():
    expected_value = 3

    strategy = PartOne()
    result = strategy.execute(str(1122))

    assert next(result) == expected_value


def test_1111():
    expected_value = 4

    strategy = PartOne()
    result = strategy.execute(str(1111))

    assert next(result) == expected_value


def test_1234():
    expected_value = 0

    strategy = PartOne()
    result = strategy.execute(str(1234))

    assert next(result) == expected_value


def test_91212129():
    expected_value = 9

    strategy = PartOne()
    result = strategy.execute(str(91212129))

    assert next(result) == expected_value
