from collections import Counter


def process_part_1(input: str) -> int:
    l1 = []
    l2 = []
    for line in input.split("\n"):
        numbers = line.split("   ")
        if numbers.__len__() == 2:
            l1.append(int(numbers[0]))
            l2.append(int(numbers[1]))
    l1.sort()
    l2.sort()
    return sum([abs(i1 - i2) for (i1, i2) in zip(l1, l2)])


def process_part_2(input: str) -> int:
    l1 = []
    l2 = []
    for line in input.split("\n"):
        numbers = line.split("   ")
        if numbers.__len__() == 2:
            l1.append(int(numbers[0]))
            l2.append(int(numbers[1]))
    map = Counter(l2)
    return sum([i * map.get(i, 0) for i in l1])
