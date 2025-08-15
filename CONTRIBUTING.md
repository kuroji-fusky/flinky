# Contributing to Flinky

## Prerequisites

To get this project running, you'll need the following:

- Docker
- Go v1.23.7 or higher
- Node 22, LTS, or higher
- PNPM package manager
  - If you don't have PNPM installed, run:

    ```console
    npm i -g pnpm
    ```
    
    You'll immediately have it installed when you run the `pnpm` command.

## Installation

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
