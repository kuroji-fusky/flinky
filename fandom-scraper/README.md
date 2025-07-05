# fandom-scraper

This was created as a quick reference on how tedious it is to curate a global furry database.

It's a scraper written a crawls Heroes and Villains Wikis to crawl the `Animals` category only to narrow the dataset for the purpose of Flinky.

## Arguments

To prevent constant requests from `.fandom.com`, a caching feature, and a `--timeout`, and `--delay` arguments are imposed to prevent `429 Too Many Requests` and `403 Forbidden` responses.

- It tracks all the URLs in a `.txt` to prevent any redundant requests