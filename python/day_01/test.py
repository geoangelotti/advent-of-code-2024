import pytest
from lib import process_part_1, process_part_2

INPUT = '''3   4
4   3
2   5
1   3
3   9
3   3
'''


@pytest.mark.parametrize("input,expected", [
    (INPUT, 11)
])
def test_part_1(input: str, expected: int):
    assert process_part_1(input) == expected


@pytest.mark.parametrize("input,expected", [
    (INPUT, 31)
])
def test_part_2(input: str, expected: int):
    assert process_part_2(input) == expected
