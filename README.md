<h1 align="center">Flinky</h1>

<p align="center">A community-curated database for all things anthro!</p>

## Premise

The main purpose of Flinky is to catalog all the fictional critters from comics, TV shows, films, series, etc. This includes, but not limited to:

- Mascots (corporate or otherwise)
- Side characters
- Characters that we're canceled or rejected for any reason

However, Flinky is **NOT** for OCs such as furry VTubers, YouTubers, etc., as there are [stringent guidelines](/client/src/content/guidelines.mdx) of what goes in and what doesn't.

## Project structure

This repo have its directories named accordingly:

- `client`: the website written in Astro
- [`server`](/server): a Go server powered using the Echo framework (as a proof-of-concept) and a web scraper in [colly](https://github.com/gocolly/colly)
- `family-tree-ref`: pseudocode reference of a soon-to-be family tree feature

### Stack

It utilizes Redis for caching of data, MeiliSearch for its search engine, and MinIO for development-only object storage mainly for storing images.

## Development

### Prerequisites

To get this project running, you'll need the following:

- Docker
- Go v1.23.7 or higher
- Node 22, LTS, or higher
- PNPM package manager
  - If you don't have PNPM installed, run `npm i -g pnpm`, you'll immediately have it installed when you run the `pnpm` command.

### Installation

> [!NOTE]
> Installing the dependencies can work the other way around. (i.e. installing the server dependencies first, then frontend, vice versa)

Install the projects' dependencies:

- From `client`

  ```console
  pnpm install
  ```
  Run the local dev server with:
  ```console
  pnpm run dev
  ```

- From `server`

  ```console
  go get download
  ```
  Run the server with:
  ```console
  go run cmd/server/server.go
  ```

## License

Copyright © 2025 Kerby Keith Aquino. Flinky is licensed under [GNU GPLv3](/LICENSE).
