<h1 align="center">Flinky</h1>

<p align="center">A community-curated database for all things anthro!</p>

---

The main purpose of Flinky is to catalog all the fictional critters from comics, games, TV shows, films, series, etc. that are *strictly* feature antropomorphic characters. This includes, but not limited to:

- Mascots (corporate or otherwise)
- Side characters
- Characters that were (or once) canceled, or rejected for any reason

However, Flinky is **NOT** for OCs as there are [stringent guidelines](/client/src/content/guidelines.mdx) of what goes in and what doesn't. This includes, but not limited to:

- Personal avatars from furry VTubers, streamers, YouTubers (i.e. Majira Strawberry, BetaEtaDelota, etc.)<sup><a href="/client/src/content/guidelines.mdx">[clause 3]</a></sup>
- Characters that claim to be from a big network but no actual coverages or press releases to back it up (i.e. AlexTheFox2000's "Toonlandia")<sup><a href="/client/src/content/guidelines.mdx">[clause 2b]</a></sup>

## Project structure

This repo have its directories named accordingly:

- [`client`](/client/): the website written in Astro
- [`server`](/server): a Go server powered using the Echo framework (as a proof-of-concept) and a web scraper in [colly](https://github.com/gocolly/colly)

### Stack

It utilizes Redis for caching of data, MeiliSearch for its search engine, and MinIO for development-only object storage mainly for storing images.

## Contributing

For contributing changes to Flinky, see [CONTRIBUTING.md](/CONTRIBUTING.md) for more details.

## License

Copyright © 2025 Fusky Labs Software. Flinky is licensed under [GNU GPLv3](/LICENSE).
