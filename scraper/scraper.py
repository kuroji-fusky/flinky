from crawl_utils import *
from constants import *

from argparse import ArgumentParser
import json
import time
import os

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

GLOBAL_TIMEOUT: int = args.timeout
GLOBAL_DELAY: int = args.delay


def _delay(sleep_sec):
    print(f"Sleeping for {sleep_sec} seconds")

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


# Create the cache file and add initial timestamps to keep track of any new or deleted entries
# the articles themselves are going to be cached
cached_last_time = 0

# TODO: Store all the articles here from cache, if available
heroes_articles: list[dict[str, str]] = []


class ScraperConfig:
    whitelist: list[str]
    ignore_filters: list[tuple[str, str]]


def main():
    # Load order:
    #   - Cached links
    #   - scraper-config.yml
    #   - If `--use-db` is true, make a connection, if it fails, ask to continue to fallback, retry, or cancel
    #   - Scrape for Heroes and Villains wikis

    # TODO: wrap this from a class
    try:
        with open("fandom_cache.txt", "r", encoding="utf-8") as f:
            print(f.readlines())

    except FileNotFoundError:
        print("Creating cache file...")
        current_time = round(time.time())

        make_file("fandom_cache.txt", f"lastcheck {current_time}\n")

    # Read the config file
    with open("scraper-config.yml", "r") as f:
        import yaml
        _scraper_config = yaml.load(f, Loader=yaml.Loader)

    ignore_filters_parsed = [list(dict.items(a))[0]
                             for a in _scraper_config["ignore-filters"]]

    scraper_config = ScraperConfig()
    scraper_config.whitelist = _scraper_config["whitelist"]
    scraper_config.ignore_filters = ignore_filters_parsed

    print(f"{len(scraper_config.whitelist)} items found in the whitelist")

    _delay(GLOBAL_DELAY)
    initial_req_url = f"https://{HEROES_BASE_URL}"

    h_page_initial = req_soup(f"{initial_req_url}/wiki/Category:Animals")  # noqa

    h_article_el = h_page_initial.select(".category-page__members-for-char a.category-page__member-link")  # noqa

    h_article_links = SoupMe.extract_links(h_article_el, prefix=initial_req_url)  # noqa
    h_article_text = SoupMe.extract_text_content(h_article_el)

    # Set the "Last" button URL, if its a match, get the remaining items, then start scraping the articles found
    end_link_marker = SoupMe.extract_links(
        h_page_initial.select_one(".category-page__pagination a:last-child")
    )

    print(end_link_marker)

    for link, text in zip(h_article_links, h_article_text):
        heroes_articles.append({
            "link": link,
            "text": text,
        })

    # print(json.dumps(heroes_articles, indent=2))


if __name__ == "__main__":
    main()
