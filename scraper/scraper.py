import sys
from crawl_utils import *
from constants import *
from argparse import ArgumentParser
import json

parser = ArgumentParser()

parser.add_argument("-db", "--use-db", action="store_true",
                    help="Enables from all data")

parser.add_argument("-D", "--delay", type=int, default=6,
                    help="Delay each request, default is 6 seconds")

parser.add_argument("-t", "--timeout", type=int, default=8,
                    help="Define the timeout of each request, default is 15 seconds")

parser.add_argument("-r", "--retries", type=int, default=8,
                    help="Retries the request once a request timeouts, default is 8 tries before the script quits prematurely")


args = parser.parse_args()


def make_file(filename: str, contents, *, json_mode=False):
    """Creates and appends content from a file for caching

    Args:
        filename (str): Your file name must have prefixes so you don't wanna die
        contents (_type_): yes
        json_mode (bool, optional): Enable JSON dumping
    """
    with open(filename, "a+", encoding="utf-8") as f:
        if not json_mode:
            f.write(contents)
        else:
            json.dump(contents, f, ensure_ascii=True, indent=2)


def main():
    if args.use_db:
        print("using db")


if __name__ == "__main__":
    if len(sys.argv) == 1:
        parser.print_help()
        sys.exit(2)
    else:
        main()
