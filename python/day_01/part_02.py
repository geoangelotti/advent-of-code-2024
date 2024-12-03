#!/usr/bin/env python

from lib import process_part_2


def main():
    with open("input.txt", "r") as file:
        print(process_part_2(file.read()))


if __name__ == "__main__":
    main()
