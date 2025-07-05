import sys
from crawl_utils import *
from argparse import ArgumentParser

parser = ArgumentParser()

parser.add_argument("--delay", type=int, default=6)
parser.add_argument("--timeout", type=int, default=8)

args = parser.parse_args()


def main():
    pass


if __name__ == "__main__":
    if len(sys.argv) == 1:
        parser.print_help()
    else:
        main()
