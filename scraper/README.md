# Flinky Scraper

A Python scraper written a crawls Heroes and Villains Wikis to crawl the `Animals` category only to narrow the dataset for the purpose of Flinky - created as a quick reference on how tedious it is to curate a global furry database.

> [!NOTE]
> To prevent constant requests from `*.fandom.com`, a caching feature, and a `--timeout`, and `--delay` arguments are imposed to prevent `4XX` responses.

The script tracks all the URLs in `fandom_cache.txt` to prevent any redundant requests and keeps a record of them; then tracks the progress with `fandom_cache.checkpoint.txt` for a worse case scenario if the script unexpectedly exits in a case of an abrupt shutdown, kernel panics

## Config

The script can be configured with `scraper-config.yml` where you can define whitelists and ignore certain URLs and strings.

## Arguments

> [!NOTE]
> While `--use-db` can technically hook up to any database, you shouldn't hook it to a production database as it only outputs incomplete metadata and the script is mainly for testing grounds 

- `--use-db`: hooks a local database from Docker
- `--timeout`: Keeps on how long the connection will be, default is 15 seconds
- `--delay`: Specifies the delay of each request, default value is 6 seconds
- `--retries`: A number of retries if a request fails before failing if the number of retries exceeded, default value is 8`