# Flinky Server

Flinky's backend written in Go using Echo framework and go-colly for scraping.

## Scraping data

> [!NOTE]
> To prevent constant requests from `*.fandom.com`, a caching feature is implemented to countermeasure this.

To speed up the development of populating the database, a scraper is necessary for this project.

The main scraper script belongs in [`scraper.go`](cmd/fandom-scraper/scraper.go) and can be fine-tuned with its config file. It scrapes certain Fandom wiki sites, and, because each wiki is different, it belongs from the `extractors` directory and it crawls these wikis:

| Wiki name     | URL                                                | Implemented | Extractor                                                     |
| ------------- | :------------------------------------------------- | ----------- | ------------------------------------------------------------- |
| Heroes Wiki   | [heroes.fandom.com](https://heroes.fandom.com)     | WIP         | [`heroes-villains.go`](scraper/extractors/heroes-villains.go) |
| Villains Wiki | [villains.fandom.com](https://villains.fandom.com) | WIP         | [`heroes-villains.go`](scraper/extractors/heroes-villains.go) |
