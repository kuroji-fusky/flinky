from crawl_utils import *
from constants import *
from argparse import ArgumentParser
import json
import time

parser = ArgumentParser()

parser.add_argument("-db", "--use-db", action="store_true",
                    help="Enables from all data")

parser.add_argument("-D", "--delay", type=int, default=6,
                    help="Delay each request, default is 6 seconds")

parser.add_argument("-t", "--timeout", type=int, default=8,
                    help="Define the timeout of each request, default is 15 seconds")

parser.add_argument("-r", "--retries", type=int, default=8,
                    help="Retries the request once a request timeouts, default is 8 tries before the script quits prematurely")

parser.add_argument("-v", "--verbose",  action="store_true",
                    help="Enables verbose logging")


args = parser.parse_args()


def _delay(sleep_sec):
    if args.verbose:
        print(f"[verbose] Sleeping for {sleep_sec} seconds")

    time.sleep(sleep_sec)


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


GLOBAL_TIMEOUT: int = args.timeout
GLOBAL_DELAY: int = args.delay


def main():
    # Create the cache file and add initial timestamps to keep track of any new or deleted entries
    # the articles themselves are going to be cached
    cached_last_time = 0
    
    # TODO: Store all the articles here from cache, if available
    heroes_articles = []

    # TODO: wrap this from a class
    try:
        with open("fandom_cache.txt", "r", encoding="utf-8") as f:
            f.readline()
    except FileNotFoundError:
        current_time = round(time.time())

        make_file("fandom_cache.txt", f"lastcheck {current_time}")

    _delay(GLOBAL_DELAY)

    # After all the checks, then we can start some scraping~
    # The "Last" button from the page, will be appended
    ending_marker = ""
    
    
    initial_req_url = f"https://{HEROES_BASE_URL}/wiki/Category:Animals"

    h_page_initial = req_soup(initial_req_url)  # noqa

    h_article_item_el = h_page_initial.select(".category-page__members-for-char a.category-page__member-link")  # noqa
    h_article_item_el = SoupMe.extract_links(h_article_item_el, prefix=initial_req_url)  # noqa

    for article in h_article_item_el:
        print(article)


if __name__ == "__main__":
    main()
